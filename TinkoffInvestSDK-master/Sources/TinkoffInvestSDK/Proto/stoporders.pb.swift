// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: stoporders.proto
//
// For information on using the generated types, please see the documentation:
//   https://github.com/apple/swift-protobuf/

import Foundation
import SwiftProtobuf

// If the compiler emits an error on this type, it is because this file
// was generated by a version of the `protoc` Swift plug-in that is
// incompatible with the version of SwiftProtobuf to which you are linking.
// Please ensure that you are building against the same version of the API
// that was used to generate this file.
fileprivate struct _GeneratedWithProtocGenSwiftVersion: SwiftProtobuf.ProtobufAPIVersionCheck {
  struct _2: SwiftProtobuf.ProtobufAPIVersion_2 {}
  typealias Version = _2
}

///Направление сделки стоп-заявки.
public enum Tinkoff_Public_Invest_Api_Contract_V1_StopOrderDirection: SwiftProtobuf.Enum {
  public typealias RawValue = Int

  ///Значение не указано.
  case unspecified // = 0

  ///Покупка.
  case buy // = 1

  ///Продажа.
  case sell // = 2
  case UNRECOGNIZED(Int)

  public init() {
    self = .unspecified
  }

  public init?(rawValue: Int) {
    switch rawValue {
    case 0: self = .unspecified
    case 1: self = .buy
    case 2: self = .sell
    default: self = .UNRECOGNIZED(rawValue)
    }
  }

  public var rawValue: Int {
    switch self {
    case .unspecified: return 0
    case .buy: return 1
    case .sell: return 2
    case .UNRECOGNIZED(let i): return i
    }
  }

}

#if swift(>=4.2)

extension Tinkoff_Public_Invest_Api_Contract_V1_StopOrderDirection: CaseIterable {
  // The compiler won't synthesize support with the UNRECOGNIZED case.
  public static var allCases: [Tinkoff_Public_Invest_Api_Contract_V1_StopOrderDirection] = [
    .unspecified,
    .buy,
    .sell,
  ]
}

#endif  // swift(>=4.2)

///Тип экспирации стоп-заявке.
public enum Tinkoff_Public_Invest_Api_Contract_V1_StopOrderExpirationType: SwiftProtobuf.Enum {
  public typealias RawValue = Int

  ///Значение не указано.
  case unspecified // = 0

  ///Действительно до отмены.
  case goodTillCancel // = 1

  ///Действительно до даты снятия.
  case goodTillDate // = 2
  case UNRECOGNIZED(Int)

  public init() {
    self = .unspecified
  }

  public init?(rawValue: Int) {
    switch rawValue {
    case 0: self = .unspecified
    case 1: self = .goodTillCancel
    case 2: self = .goodTillDate
    default: self = .UNRECOGNIZED(rawValue)
    }
  }

  public var rawValue: Int {
    switch self {
    case .unspecified: return 0
    case .goodTillCancel: return 1
    case .goodTillDate: return 2
    case .UNRECOGNIZED(let i): return i
    }
  }

}

#if swift(>=4.2)

extension Tinkoff_Public_Invest_Api_Contract_V1_StopOrderExpirationType: CaseIterable {
  // The compiler won't synthesize support with the UNRECOGNIZED case.
  public static var allCases: [Tinkoff_Public_Invest_Api_Contract_V1_StopOrderExpirationType] = [
    .unspecified,
    .goodTillCancel,
    .goodTillDate,
  ]
}

#endif  // swift(>=4.2)

///Тип стоп-заявки.
public enum Tinkoff_Public_Invest_Api_Contract_V1_StopOrderType: SwiftProtobuf.Enum {
  public typealias RawValue = Int

  ///Значение не указано.
  case unspecified // = 0

  ///Take-profit заявка.
  case takeProfit // = 1

  ///Stop-loss заявка.
  case stopLoss // = 2

  ///Stop-limit заявка.
  case stopLimit // = 3
  case UNRECOGNIZED(Int)

  public init() {
    self = .unspecified
  }

  public init?(rawValue: Int) {
    switch rawValue {
    case 0: self = .unspecified
    case 1: self = .takeProfit
    case 2: self = .stopLoss
    case 3: self = .stopLimit
    default: self = .UNRECOGNIZED(rawValue)
    }
  }

  public var rawValue: Int {
    switch self {
    case .unspecified: return 0
    case .takeProfit: return 1
    case .stopLoss: return 2
    case .stopLimit: return 3
    case .UNRECOGNIZED(let i): return i
    }
  }

}

#if swift(>=4.2)

extension Tinkoff_Public_Invest_Api_Contract_V1_StopOrderType: CaseIterable {
  // The compiler won't synthesize support with the UNRECOGNIZED case.
  public static var allCases: [Tinkoff_Public_Invest_Api_Contract_V1_StopOrderType] = [
    .unspecified,
    .takeProfit,
    .stopLoss,
    .stopLimit,
  ]
}

#endif  // swift(>=4.2)

///Запрос выставления стоп-заявки.
public struct Tinkoff_Public_Invest_Api_Contract_V1_PostStopOrderRequest {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///Figi-идентификатор инструмента.
  public var figi: String = String()

  ///Количество лотов.
  public var quantity: Int64 = 0

  ///Цена за 1 инструмент. Для получения стоимости лота требуется умножить на лотность инструмента.
  public var price: Tinkoff_Public_Invest_Api_Contract_V1_Quotation {
    get {return _price ?? Tinkoff_Public_Invest_Api_Contract_V1_Quotation()}
    set {_price = newValue}
  }
  /// Returns true if `price` has been explicitly set.
  public var hasPrice: Bool {return self._price != nil}
  /// Clears the value of `price`. Subsequent reads from it will return its default value.
  public mutating func clearPrice() {self._price = nil}

  ///Стоп-цена заявки за 1 инструмент. Для получения стоимости лота требуется умножить на лотность инструмента.
  public var stopPrice: Tinkoff_Public_Invest_Api_Contract_V1_Quotation {
    get {return _stopPrice ?? Tinkoff_Public_Invest_Api_Contract_V1_Quotation()}
    set {_stopPrice = newValue}
  }
  /// Returns true if `stopPrice` has been explicitly set.
  public var hasStopPrice: Bool {return self._stopPrice != nil}
  /// Clears the value of `stopPrice`. Subsequent reads from it will return its default value.
  public mutating func clearStopPrice() {self._stopPrice = nil}

  ///Направление операции.
  public var direction: Tinkoff_Public_Invest_Api_Contract_V1_StopOrderDirection = .unspecified

  ///Номер счёта.
  public var accountID: String = String()

  ///Тип экспирации заявки.
  public var expirationType: Tinkoff_Public_Invest_Api_Contract_V1_StopOrderExpirationType = .unspecified

  ///Тип заявки.
  public var stopOrderType: Tinkoff_Public_Invest_Api_Contract_V1_StopOrderType = .unspecified

  ///Дата и время окончания действия стоп-заявки в часовом поясе UTC. **Для ExpirationType = GoodTillDate заполнение обязательно**.
  public var expireDate: SwiftProtobuf.Google_Protobuf_Timestamp {
    get {return _expireDate ?? SwiftProtobuf.Google_Protobuf_Timestamp()}
    set {_expireDate = newValue}
  }
  /// Returns true if `expireDate` has been explicitly set.
  public var hasExpireDate: Bool {return self._expireDate != nil}
  /// Clears the value of `expireDate`. Subsequent reads from it will return its default value.
  public mutating func clearExpireDate() {self._expireDate = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _price: Tinkoff_Public_Invest_Api_Contract_V1_Quotation? = nil
  fileprivate var _stopPrice: Tinkoff_Public_Invest_Api_Contract_V1_Quotation? = nil
  fileprivate var _expireDate: SwiftProtobuf.Google_Protobuf_Timestamp? = nil
}

///Результат выставления стоп-заявки.
public struct Tinkoff_Public_Invest_Api_Contract_V1_PostStopOrderResponse {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///Уникальный идентификатор стоп-заявки.
  public var stopOrderID: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

///Запрос получения списка активных стоп-заявок.
public struct Tinkoff_Public_Invest_Api_Contract_V1_GetStopOrdersRequest {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///Идентификатор счёта клиента.
  public var accountID: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

///Список активных стоп-заявок.
public struct Tinkoff_Public_Invest_Api_Contract_V1_GetStopOrdersResponse {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///Массив стоп-заявок по счёту.
  public var stopOrders: [Tinkoff_Public_Invest_Api_Contract_V1_StopOrder] = []

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

///Запрос отмены выставленной стоп-заявки.
public struct Tinkoff_Public_Invest_Api_Contract_V1_CancelStopOrderRequest {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///Идентификатор счёта клиента.
  public var accountID: String = String()

  ///Уникальный идентификатор стоп-заявки.
  public var stopOrderID: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

///Результат отмены выставленной стоп-заявки.
public struct Tinkoff_Public_Invest_Api_Contract_V1_CancelStopOrderResponse {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///Время отмены заявки в часовом поясе UTC.
  public var time: SwiftProtobuf.Google_Protobuf_Timestamp {
    get {return _time ?? SwiftProtobuf.Google_Protobuf_Timestamp()}
    set {_time = newValue}
  }
  /// Returns true if `time` has been explicitly set.
  public var hasTime: Bool {return self._time != nil}
  /// Clears the value of `time`. Subsequent reads from it will return its default value.
  public mutating func clearTime() {self._time = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _time: SwiftProtobuf.Google_Protobuf_Timestamp? = nil
}

///Информация о стоп-заявке.
public struct Tinkoff_Public_Invest_Api_Contract_V1_StopOrder {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///Идентификатор-идентификатор стоп-заявки.
  public var stopOrderID: String {
    get {return _storage._stopOrderID}
    set {_uniqueStorage()._stopOrderID = newValue}
  }

  ///Запрошено лотов.
  public var lotsRequested: Int64 {
    get {return _storage._lotsRequested}
    set {_uniqueStorage()._lotsRequested = newValue}
  }

  ///Figi-идентификатор инструмента.
  public var figi: String {
    get {return _storage._figi}
    set {_uniqueStorage()._figi = newValue}
  }

  ///Направление операции.
  public var direction: Tinkoff_Public_Invest_Api_Contract_V1_StopOrderDirection {
    get {return _storage._direction}
    set {_uniqueStorage()._direction = newValue}
  }

  ///Валюта стоп-заявки.
  public var currency: String {
    get {return _storage._currency}
    set {_uniqueStorage()._currency = newValue}
  }

  ///Тип стоп-заявки.
  public var orderType: Tinkoff_Public_Invest_Api_Contract_V1_StopOrderType {
    get {return _storage._orderType}
    set {_uniqueStorage()._orderType = newValue}
  }

  ///Дата и время выставления заявки в часовом поясе UTC.
  public var createDate: SwiftProtobuf.Google_Protobuf_Timestamp {
    get {return _storage._createDate ?? SwiftProtobuf.Google_Protobuf_Timestamp()}
    set {_uniqueStorage()._createDate = newValue}
  }
  /// Returns true if `createDate` has been explicitly set.
  public var hasCreateDate: Bool {return _storage._createDate != nil}
  /// Clears the value of `createDate`. Subsequent reads from it will return its default value.
  public mutating func clearCreateDate() {_uniqueStorage()._createDate = nil}

  ///Дата и время конвертации стоп-заявки в биржевую в часовом поясе UTC.
  public var activationDateTime: SwiftProtobuf.Google_Protobuf_Timestamp {
    get {return _storage._activationDateTime ?? SwiftProtobuf.Google_Protobuf_Timestamp()}
    set {_uniqueStorage()._activationDateTime = newValue}
  }
  /// Returns true if `activationDateTime` has been explicitly set.
  public var hasActivationDateTime: Bool {return _storage._activationDateTime != nil}
  /// Clears the value of `activationDateTime`. Subsequent reads from it will return its default value.
  public mutating func clearActivationDateTime() {_uniqueStorage()._activationDateTime = nil}

  ///Дата и время снятия заявки в часовом поясе UTC.
  public var expirationTime: SwiftProtobuf.Google_Protobuf_Timestamp {
    get {return _storage._expirationTime ?? SwiftProtobuf.Google_Protobuf_Timestamp()}
    set {_uniqueStorage()._expirationTime = newValue}
  }
  /// Returns true if `expirationTime` has been explicitly set.
  public var hasExpirationTime: Bool {return _storage._expirationTime != nil}
  /// Clears the value of `expirationTime`. Subsequent reads from it will return its default value.
  public mutating func clearExpirationTime() {_uniqueStorage()._expirationTime = nil}

  ///Цена заявки за 1 инструмент. Для получения стоимости лота требуется умножить на лотность инструмента.
  public var price: Tinkoff_Public_Invest_Api_Contract_V1_MoneyValue {
    get {return _storage._price ?? Tinkoff_Public_Invest_Api_Contract_V1_MoneyValue()}
    set {_uniqueStorage()._price = newValue}
  }
  /// Returns true if `price` has been explicitly set.
  public var hasPrice: Bool {return _storage._price != nil}
  /// Clears the value of `price`. Subsequent reads from it will return its default value.
  public mutating func clearPrice() {_uniqueStorage()._price = nil}

  ///Цена активации стоп-заявки за 1 инструмент. Для получения стоимости лота требуется умножить на лотность инструмента.
  public var stopPrice: Tinkoff_Public_Invest_Api_Contract_V1_MoneyValue {
    get {return _storage._stopPrice ?? Tinkoff_Public_Invest_Api_Contract_V1_MoneyValue()}
    set {_uniqueStorage()._stopPrice = newValue}
  }
  /// Returns true if `stopPrice` has been explicitly set.
  public var hasStopPrice: Bool {return _storage._stopPrice != nil}
  /// Clears the value of `stopPrice`. Subsequent reads from it will return its default value.
  public mutating func clearStopPrice() {_uniqueStorage()._stopPrice = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _storage = _StorageClass.defaultInstance
}

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "tinkoff.public.invest.api.contract.v1"

extension Tinkoff_Public_Invest_Api_Contract_V1_StopOrderDirection: SwiftProtobuf._ProtoNameProviding {
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "STOP_ORDER_DIRECTION_UNSPECIFIED"),
    1: .same(proto: "STOP_ORDER_DIRECTION_BUY"),
    2: .same(proto: "STOP_ORDER_DIRECTION_SELL"),
  ]
}

extension Tinkoff_Public_Invest_Api_Contract_V1_StopOrderExpirationType: SwiftProtobuf._ProtoNameProviding {
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "STOP_ORDER_EXPIRATION_TYPE_UNSPECIFIED"),
    1: .same(proto: "STOP_ORDER_EXPIRATION_TYPE_GOOD_TILL_CANCEL"),
    2: .same(proto: "STOP_ORDER_EXPIRATION_TYPE_GOOD_TILL_DATE"),
  ]
}

extension Tinkoff_Public_Invest_Api_Contract_V1_StopOrderType: SwiftProtobuf._ProtoNameProviding {
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "STOP_ORDER_TYPE_UNSPECIFIED"),
    1: .same(proto: "STOP_ORDER_TYPE_TAKE_PROFIT"),
    2: .same(proto: "STOP_ORDER_TYPE_STOP_LOSS"),
    3: .same(proto: "STOP_ORDER_TYPE_STOP_LIMIT"),
  ]
}

extension Tinkoff_Public_Invest_Api_Contract_V1_PostStopOrderRequest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".PostStopOrderRequest"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "figi"),
    2: .same(proto: "quantity"),
    3: .same(proto: "price"),
    4: .standard(proto: "stop_price"),
    5: .same(proto: "direction"),
    6: .standard(proto: "account_id"),
    7: .standard(proto: "expiration_type"),
    8: .standard(proto: "stop_order_type"),
    9: .standard(proto: "expire_date"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.figi) }()
      case 2: try { try decoder.decodeSingularInt64Field(value: &self.quantity) }()
      case 3: try { try decoder.decodeSingularMessageField(value: &self._price) }()
      case 4: try { try decoder.decodeSingularMessageField(value: &self._stopPrice) }()
      case 5: try { try decoder.decodeSingularEnumField(value: &self.direction) }()
      case 6: try { try decoder.decodeSingularStringField(value: &self.accountID) }()
      case 7: try { try decoder.decodeSingularEnumField(value: &self.expirationType) }()
      case 8: try { try decoder.decodeSingularEnumField(value: &self.stopOrderType) }()
      case 9: try { try decoder.decodeSingularMessageField(value: &self._expireDate) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    if !self.figi.isEmpty {
      try visitor.visitSingularStringField(value: self.figi, fieldNumber: 1)
    }
    if self.quantity != 0 {
      try visitor.visitSingularInt64Field(value: self.quantity, fieldNumber: 2)
    }
    try { if let v = self._price {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 3)
    } }()
    try { if let v = self._stopPrice {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 4)
    } }()
    if self.direction != .unspecified {
      try visitor.visitSingularEnumField(value: self.direction, fieldNumber: 5)
    }
    if !self.accountID.isEmpty {
      try visitor.visitSingularStringField(value: self.accountID, fieldNumber: 6)
    }
    if self.expirationType != .unspecified {
      try visitor.visitSingularEnumField(value: self.expirationType, fieldNumber: 7)
    }
    if self.stopOrderType != .unspecified {
      try visitor.visitSingularEnumField(value: self.stopOrderType, fieldNumber: 8)
    }
    try { if let v = self._expireDate {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 9)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Tinkoff_Public_Invest_Api_Contract_V1_PostStopOrderRequest, rhs: Tinkoff_Public_Invest_Api_Contract_V1_PostStopOrderRequest) -> Bool {
    if lhs.figi != rhs.figi {return false}
    if lhs.quantity != rhs.quantity {return false}
    if lhs._price != rhs._price {return false}
    if lhs._stopPrice != rhs._stopPrice {return false}
    if lhs.direction != rhs.direction {return false}
    if lhs.accountID != rhs.accountID {return false}
    if lhs.expirationType != rhs.expirationType {return false}
    if lhs.stopOrderType != rhs.stopOrderType {return false}
    if lhs._expireDate != rhs._expireDate {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Tinkoff_Public_Invest_Api_Contract_V1_PostStopOrderResponse: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".PostStopOrderResponse"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "stop_order_id"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.stopOrderID) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.stopOrderID.isEmpty {
      try visitor.visitSingularStringField(value: self.stopOrderID, fieldNumber: 1)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Tinkoff_Public_Invest_Api_Contract_V1_PostStopOrderResponse, rhs: Tinkoff_Public_Invest_Api_Contract_V1_PostStopOrderResponse) -> Bool {
    if lhs.stopOrderID != rhs.stopOrderID {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Tinkoff_Public_Invest_Api_Contract_V1_GetStopOrdersRequest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".GetStopOrdersRequest"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "account_id"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.accountID) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.accountID.isEmpty {
      try visitor.visitSingularStringField(value: self.accountID, fieldNumber: 1)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Tinkoff_Public_Invest_Api_Contract_V1_GetStopOrdersRequest, rhs: Tinkoff_Public_Invest_Api_Contract_V1_GetStopOrdersRequest) -> Bool {
    if lhs.accountID != rhs.accountID {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Tinkoff_Public_Invest_Api_Contract_V1_GetStopOrdersResponse: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".GetStopOrdersResponse"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "stop_orders"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeRepeatedMessageField(value: &self.stopOrders) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.stopOrders.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.stopOrders, fieldNumber: 1)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Tinkoff_Public_Invest_Api_Contract_V1_GetStopOrdersResponse, rhs: Tinkoff_Public_Invest_Api_Contract_V1_GetStopOrdersResponse) -> Bool {
    if lhs.stopOrders != rhs.stopOrders {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Tinkoff_Public_Invest_Api_Contract_V1_CancelStopOrderRequest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CancelStopOrderRequest"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "account_id"),
    2: .standard(proto: "stop_order_id"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.accountID) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.stopOrderID) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.accountID.isEmpty {
      try visitor.visitSingularStringField(value: self.accountID, fieldNumber: 1)
    }
    if !self.stopOrderID.isEmpty {
      try visitor.visitSingularStringField(value: self.stopOrderID, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Tinkoff_Public_Invest_Api_Contract_V1_CancelStopOrderRequest, rhs: Tinkoff_Public_Invest_Api_Contract_V1_CancelStopOrderRequest) -> Bool {
    if lhs.accountID != rhs.accountID {return false}
    if lhs.stopOrderID != rhs.stopOrderID {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Tinkoff_Public_Invest_Api_Contract_V1_CancelStopOrderResponse: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CancelStopOrderResponse"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "time"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._time) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._time {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Tinkoff_Public_Invest_Api_Contract_V1_CancelStopOrderResponse, rhs: Tinkoff_Public_Invest_Api_Contract_V1_CancelStopOrderResponse) -> Bool {
    if lhs._time != rhs._time {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Tinkoff_Public_Invest_Api_Contract_V1_StopOrder: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".StopOrder"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "stop_order_id"),
    2: .standard(proto: "lots_requested"),
    3: .same(proto: "figi"),
    4: .same(proto: "direction"),
    5: .same(proto: "currency"),
    6: .standard(proto: "order_type"),
    7: .standard(proto: "create_date"),
    8: .standard(proto: "activation_date_time"),
    9: .standard(proto: "expiration_time"),
    10: .same(proto: "price"),
    11: .standard(proto: "stop_price"),
  ]

  fileprivate class _StorageClass {
    var _stopOrderID: String = String()
    var _lotsRequested: Int64 = 0
    var _figi: String = String()
    var _direction: Tinkoff_Public_Invest_Api_Contract_V1_StopOrderDirection = .unspecified
    var _currency: String = String()
    var _orderType: Tinkoff_Public_Invest_Api_Contract_V1_StopOrderType = .unspecified
    var _createDate: SwiftProtobuf.Google_Protobuf_Timestamp? = nil
    var _activationDateTime: SwiftProtobuf.Google_Protobuf_Timestamp? = nil
    var _expirationTime: SwiftProtobuf.Google_Protobuf_Timestamp? = nil
    var _price: Tinkoff_Public_Invest_Api_Contract_V1_MoneyValue? = nil
    var _stopPrice: Tinkoff_Public_Invest_Api_Contract_V1_MoneyValue? = nil

    static let defaultInstance = _StorageClass()

    private init() {}

    init(copying source: _StorageClass) {
      _stopOrderID = source._stopOrderID
      _lotsRequested = source._lotsRequested
      _figi = source._figi
      _direction = source._direction
      _currency = source._currency
      _orderType = source._orderType
      _createDate = source._createDate
      _activationDateTime = source._activationDateTime
      _expirationTime = source._expirationTime
      _price = source._price
      _stopPrice = source._stopPrice
    }
  }

  fileprivate mutating func _uniqueStorage() -> _StorageClass {
    if !isKnownUniquelyReferenced(&_storage) {
      _storage = _StorageClass(copying: _storage)
    }
    return _storage
  }

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    _ = _uniqueStorage()
    try withExtendedLifetime(_storage) { (_storage: _StorageClass) in
      while let fieldNumber = try decoder.nextFieldNumber() {
        // The use of inline closures is to circumvent an issue where the compiler
        // allocates stack space for every case branch when no optimizations are
        // enabled. https://github.com/apple/swift-protobuf/issues/1034
        switch fieldNumber {
        case 1: try { try decoder.decodeSingularStringField(value: &_storage._stopOrderID) }()
        case 2: try { try decoder.decodeSingularInt64Field(value: &_storage._lotsRequested) }()
        case 3: try { try decoder.decodeSingularStringField(value: &_storage._figi) }()
        case 4: try { try decoder.decodeSingularEnumField(value: &_storage._direction) }()
        case 5: try { try decoder.decodeSingularStringField(value: &_storage._currency) }()
        case 6: try { try decoder.decodeSingularEnumField(value: &_storage._orderType) }()
        case 7: try { try decoder.decodeSingularMessageField(value: &_storage._createDate) }()
        case 8: try { try decoder.decodeSingularMessageField(value: &_storage._activationDateTime) }()
        case 9: try { try decoder.decodeSingularMessageField(value: &_storage._expirationTime) }()
        case 10: try { try decoder.decodeSingularMessageField(value: &_storage._price) }()
        case 11: try { try decoder.decodeSingularMessageField(value: &_storage._stopPrice) }()
        default: break
        }
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    try withExtendedLifetime(_storage) { (_storage: _StorageClass) in
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every if/case branch local when no optimizations
      // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
      // https://github.com/apple/swift-protobuf/issues/1182
      if !_storage._stopOrderID.isEmpty {
        try visitor.visitSingularStringField(value: _storage._stopOrderID, fieldNumber: 1)
      }
      if _storage._lotsRequested != 0 {
        try visitor.visitSingularInt64Field(value: _storage._lotsRequested, fieldNumber: 2)
      }
      if !_storage._figi.isEmpty {
        try visitor.visitSingularStringField(value: _storage._figi, fieldNumber: 3)
      }
      if _storage._direction != .unspecified {
        try visitor.visitSingularEnumField(value: _storage._direction, fieldNumber: 4)
      }
      if !_storage._currency.isEmpty {
        try visitor.visitSingularStringField(value: _storage._currency, fieldNumber: 5)
      }
      if _storage._orderType != .unspecified {
        try visitor.visitSingularEnumField(value: _storage._orderType, fieldNumber: 6)
      }
      try { if let v = _storage._createDate {
        try visitor.visitSingularMessageField(value: v, fieldNumber: 7)
      } }()
      try { if let v = _storage._activationDateTime {
        try visitor.visitSingularMessageField(value: v, fieldNumber: 8)
      } }()
      try { if let v = _storage._expirationTime {
        try visitor.visitSingularMessageField(value: v, fieldNumber: 9)
      } }()
      try { if let v = _storage._price {
        try visitor.visitSingularMessageField(value: v, fieldNumber: 10)
      } }()
      try { if let v = _storage._stopPrice {
        try visitor.visitSingularMessageField(value: v, fieldNumber: 11)
      } }()
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Tinkoff_Public_Invest_Api_Contract_V1_StopOrder, rhs: Tinkoff_Public_Invest_Api_Contract_V1_StopOrder) -> Bool {
    if lhs._storage !== rhs._storage {
      let storagesAreEqual: Bool = withExtendedLifetime((lhs._storage, rhs._storage)) { (_args: (_StorageClass, _StorageClass)) in
        let _storage = _args.0
        let rhs_storage = _args.1
        if _storage._stopOrderID != rhs_storage._stopOrderID {return false}
        if _storage._lotsRequested != rhs_storage._lotsRequested {return false}
        if _storage._figi != rhs_storage._figi {return false}
        if _storage._direction != rhs_storage._direction {return false}
        if _storage._currency != rhs_storage._currency {return false}
        if _storage._orderType != rhs_storage._orderType {return false}
        if _storage._createDate != rhs_storage._createDate {return false}
        if _storage._activationDateTime != rhs_storage._activationDateTime {return false}
        if _storage._expirationTime != rhs_storage._expirationTime {return false}
        if _storage._price != rhs_storage._price {return false}
        if _storage._stopPrice != rhs_storage._stopPrice {return false}
        return true
      }
      if !storagesAreEqual {return false}
    }
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}
