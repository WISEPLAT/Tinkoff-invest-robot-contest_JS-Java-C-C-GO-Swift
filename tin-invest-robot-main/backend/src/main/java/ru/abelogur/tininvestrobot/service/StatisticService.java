package ru.abelogur.tininvestrobot.service;

import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Service;
import ru.abelogur.tininvestrobot.controller.exception.RestRuntimeException;
import ru.abelogur.tininvestrobot.domain.InvestBot;
import ru.abelogur.tininvestrobot.domain.OrderAction;
import ru.abelogur.tininvestrobot.domain.TradeType;
import ru.abelogur.tininvestrobot.dto.StatisticDto;
import ru.abelogur.tininvestrobot.repository.CurrencyRepository;
import ru.abelogur.tininvestrobot.repository.InstrumentRepository;
import ru.abelogur.tininvestrobot.repository.InvestBotRepository;
import ru.abelogur.tininvestrobot.repository.OrderHistoryRepository;
import ru.tinkoff.piapi.core.utils.MapperUtils;

import java.math.BigDecimal;
import java.util.List;
import java.util.UUID;
import java.util.stream.Collectors;

@Service
@RequiredArgsConstructor
public class StatisticService {

    private final OrderHistoryRepository orderHistoryRepository;
    private final SdkService sdkService;
    private final InstrumentRepository instrumentRepository;
    private final InvestBotRepository botRepository;
    private final CurrencyRepository currencyRepository;

    public StatisticDto getStatistic(UUID botUuid) {
        var settings = botRepository.get(botUuid)
                .map(InvestBot::getState)
                .orElseThrow(() -> new RestRuntimeException("Бота с таким uuid не найден", HttpStatus.NOT_FOUND));
        var orders = orderHistoryRepository.getAll(botUuid);
        if (orders.isEmpty()) {
            return StatisticDto.empty();
        }

        var profit = BigDecimal.ZERO;
        var commission = BigDecimal.ZERO;
        var usedMoney = BigDecimal.ZERO;
        var longs = 0;
        var shorts = 0;

        for (var order : orders) {
            if (order.getAction().equals(OrderAction.OPEN) && order.getType().equals(TradeType.LONG)) {
                longs++;
                profit = profit.subtract(order.getPrice());
            } else if (order.getAction().equals(OrderAction.CLOSE) && order.getType().equals(TradeType.SHORT)) {
                shorts--;
                profit = profit.subtract(order.getPrice());
            } else if (order.getAction().equals(OrderAction.CLOSE) && order.getType().equals(TradeType.LONG)) {
                longs--;
                profit = profit.add(order.getPrice());
            } else if (order.getAction().equals(OrderAction.OPEN) && order.getType().equals(TradeType.SHORT)) {
                shorts++;
                profit = profit.add(order.getPrice());
            }
            commission = commission.add(order.getCommission());
            if (profit.compareTo(usedMoney.negate()) < 0) {
                usedMoney = profit.negate();
            }
        }

        var lastPrice = getLastPrice(settings.getFigi());
        profit = profit.add(lastPrice.multiply(BigDecimal.valueOf(longs)));
        profit = profit.subtract(lastPrice.multiply(BigDecimal.valueOf(shorts)));
        return new StatisticDto(orders.stream()
                .sorted(((o1, o2) -> o2.getTime().compareTo(o1.getTime())))
                .collect(Collectors.toList()), profit, commission, usedMoney);
    }

    public StatisticDto getGeneralStatistic() {
        return botRepository.getAll().stream()
                .map(bot -> {
                    var instrument = instrumentRepository.get(bot.getState().getGroupId().getFigi());
                    var currencyMultiplier = currencyRepository.getCurrentPrice(instrument.getCurrency());
                    return getStatistic(bot.getState().getUuid()).multiply(currencyMultiplier);
                })
                .reduce(StatisticDto.empty(), StatisticDto::sum);
    }

    private BigDecimal getLastPrice(String figi) {
        var price = sdkService.getInvestApi().getMarketDataService().getLastPricesSync(List.of(figi)).get(0);
        return MapperUtils.quotationToBigDecimal(price.getPrice()).multiply(instrumentRepository.get(figi).getLot());
    }
}
