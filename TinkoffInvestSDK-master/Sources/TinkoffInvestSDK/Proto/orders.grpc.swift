//
// DO NOT EDIT.
//
// Generated by the protocol buffer compiler.
// Source: orders.proto
//

//
// Copyright 2018, gRPC Authors All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
import GRPC
import NIO
import SwiftProtobuf


/// Usage: instantiate `Tinkoff_Public_Invest_Api_Contract_V1_OrdersStreamServiceClient`, then call methods of this protocol to make API calls.
public protocol Tinkoff_Public_Invest_Api_Contract_V1_OrdersStreamServiceClientProtocol: GRPCClient {
  var serviceName: String { get }
  var interceptors: Tinkoff_Public_Invest_Api_Contract_V1_OrdersStreamServiceClientInterceptorFactoryProtocol? { get }

  func tradesStream(
    _ request: Tinkoff_Public_Invest_Api_Contract_V1_TradesStreamRequest,
    callOptions: CallOptions?,
    handler: @escaping (Tinkoff_Public_Invest_Api_Contract_V1_TradesStreamResponse) -> Void
  ) -> ServerStreamingCall<Tinkoff_Public_Invest_Api_Contract_V1_TradesStreamRequest, Tinkoff_Public_Invest_Api_Contract_V1_TradesStreamResponse>
}

extension Tinkoff_Public_Invest_Api_Contract_V1_OrdersStreamServiceClientProtocol {
  public var serviceName: String {
    return "tinkoff.public.invest.api.contract.v1.OrdersStreamService"
  }

  ///Stream сделок пользователя
  ///
  /// - Parameters:
  ///   - request: Request to send to TradesStream.
  ///   - callOptions: Call options.
  ///   - handler: A closure called when each response is received from the server.
  /// - Returns: A `ServerStreamingCall` with futures for the metadata and status.
  public func tradesStream(
    _ request: Tinkoff_Public_Invest_Api_Contract_V1_TradesStreamRequest,
    callOptions: CallOptions? = nil,
    handler: @escaping (Tinkoff_Public_Invest_Api_Contract_V1_TradesStreamResponse) -> Void
  ) -> ServerStreamingCall<Tinkoff_Public_Invest_Api_Contract_V1_TradesStreamRequest, Tinkoff_Public_Invest_Api_Contract_V1_TradesStreamResponse> {
    return self.makeServerStreamingCall(
      path: "/tinkoff.public.invest.api.contract.v1.OrdersStreamService/TradesStream",
      request: request,
      callOptions: callOptions ?? self.defaultCallOptions,
      interceptors: self.interceptors?.makeTradesStreamInterceptors() ?? [],
      handler: handler
    )
  }
}

public protocol Tinkoff_Public_Invest_Api_Contract_V1_OrdersStreamServiceClientInterceptorFactoryProtocol {

  /// - Returns: Interceptors to use when invoking 'tradesStream'.
  func makeTradesStreamInterceptors() -> [ClientInterceptor<Tinkoff_Public_Invest_Api_Contract_V1_TradesStreamRequest, Tinkoff_Public_Invest_Api_Contract_V1_TradesStreamResponse>]
}

public final class Tinkoff_Public_Invest_Api_Contract_V1_OrdersStreamServiceClient: Tinkoff_Public_Invest_Api_Contract_V1_OrdersStreamServiceClientProtocol {
  public let channel: GRPCChannel
  public var defaultCallOptions: CallOptions
  public var interceptors: Tinkoff_Public_Invest_Api_Contract_V1_OrdersStreamServiceClientInterceptorFactoryProtocol?

  /// Creates a client for the tinkoff.public.invest.api.contract.v1.OrdersStreamService service.
  ///
  /// - Parameters:
  ///   - channel: `GRPCChannel` to the service host.
  ///   - defaultCallOptions: Options to use for each service call if the user doesn't provide them.
  ///   - interceptors: A factory providing interceptors for each RPC.
  public init(
    channel: GRPCChannel,
    defaultCallOptions: CallOptions = CallOptions(),
    interceptors: Tinkoff_Public_Invest_Api_Contract_V1_OrdersStreamServiceClientInterceptorFactoryProtocol? = nil
  ) {
    self.channel = channel
    self.defaultCallOptions = defaultCallOptions
    self.interceptors = interceptors
  }
}

/// Сервис предназначен для работы с торговыми поручениями:</br> **1**.
///выставление;</br> **2**. отмена;</br> **3**. получение статуса;</br> **4**.
///расчёт полной стоимости;</br> **5**. получение списка заявок.
///
/// Usage: instantiate `Tinkoff_Public_Invest_Api_Contract_V1_OrdersServiceClient`, then call methods of this protocol to make API calls.
public protocol Tinkoff_Public_Invest_Api_Contract_V1_OrdersServiceClientProtocol: GRPCClient {
  var serviceName: String { get }
  var interceptors: Tinkoff_Public_Invest_Api_Contract_V1_OrdersServiceClientInterceptorFactoryProtocol? { get }

  func postOrder(
    _ request: Tinkoff_Public_Invest_Api_Contract_V1_PostOrderRequest,
    callOptions: CallOptions?
  ) -> UnaryCall<Tinkoff_Public_Invest_Api_Contract_V1_PostOrderRequest, Tinkoff_Public_Invest_Api_Contract_V1_PostOrderResponse>

  func cancelOrder(
    _ request: Tinkoff_Public_Invest_Api_Contract_V1_CancelOrderRequest,
    callOptions: CallOptions?
  ) -> UnaryCall<Tinkoff_Public_Invest_Api_Contract_V1_CancelOrderRequest, Tinkoff_Public_Invest_Api_Contract_V1_CancelOrderResponse>

  func getOrderState(
    _ request: Tinkoff_Public_Invest_Api_Contract_V1_GetOrderStateRequest,
    callOptions: CallOptions?
  ) -> UnaryCall<Tinkoff_Public_Invest_Api_Contract_V1_GetOrderStateRequest, Tinkoff_Public_Invest_Api_Contract_V1_OrderState>

  func getOrders(
    _ request: Tinkoff_Public_Invest_Api_Contract_V1_GetOrdersRequest,
    callOptions: CallOptions?
  ) -> UnaryCall<Tinkoff_Public_Invest_Api_Contract_V1_GetOrdersRequest, Tinkoff_Public_Invest_Api_Contract_V1_GetOrdersResponse>
}

extension Tinkoff_Public_Invest_Api_Contract_V1_OrdersServiceClientProtocol {
  public var serviceName: String {
    return "tinkoff.public.invest.api.contract.v1.OrdersService"
  }

  ///Метод выставления заявки.
  ///
  /// - Parameters:
  ///   - request: Request to send to PostOrder.
  ///   - callOptions: Call options.
  /// - Returns: A `UnaryCall` with futures for the metadata, status and response.
  public func postOrder(
    _ request: Tinkoff_Public_Invest_Api_Contract_V1_PostOrderRequest,
    callOptions: CallOptions? = nil
  ) -> UnaryCall<Tinkoff_Public_Invest_Api_Contract_V1_PostOrderRequest, Tinkoff_Public_Invest_Api_Contract_V1_PostOrderResponse> {
    return self.makeUnaryCall(
      path: "/tinkoff.public.invest.api.contract.v1.OrdersService/PostOrder",
      request: request,
      callOptions: callOptions ?? self.defaultCallOptions,
      interceptors: self.interceptors?.makePostOrderInterceptors() ?? []
    )
  }

  ///Метод отмены биржевой заявки.
  ///
  /// - Parameters:
  ///   - request: Request to send to CancelOrder.
  ///   - callOptions: Call options.
  /// - Returns: A `UnaryCall` with futures for the metadata, status and response.
  public func cancelOrder(
    _ request: Tinkoff_Public_Invest_Api_Contract_V1_CancelOrderRequest,
    callOptions: CallOptions? = nil
  ) -> UnaryCall<Tinkoff_Public_Invest_Api_Contract_V1_CancelOrderRequest, Tinkoff_Public_Invest_Api_Contract_V1_CancelOrderResponse> {
    return self.makeUnaryCall(
      path: "/tinkoff.public.invest.api.contract.v1.OrdersService/CancelOrder",
      request: request,
      callOptions: callOptions ?? self.defaultCallOptions,
      interceptors: self.interceptors?.makeCancelOrderInterceptors() ?? []
    )
  }

  ///Метод получения статуса торгового поручения.
  ///
  /// - Parameters:
  ///   - request: Request to send to GetOrderState.
  ///   - callOptions: Call options.
  /// - Returns: A `UnaryCall` with futures for the metadata, status and response.
  public func getOrderState(
    _ request: Tinkoff_Public_Invest_Api_Contract_V1_GetOrderStateRequest,
    callOptions: CallOptions? = nil
  ) -> UnaryCall<Tinkoff_Public_Invest_Api_Contract_V1_GetOrderStateRequest, Tinkoff_Public_Invest_Api_Contract_V1_OrderState> {
    return self.makeUnaryCall(
      path: "/tinkoff.public.invest.api.contract.v1.OrdersService/GetOrderState",
      request: request,
      callOptions: callOptions ?? self.defaultCallOptions,
      interceptors: self.interceptors?.makeGetOrderStateInterceptors() ?? []
    )
  }

  ///Метод получения списка активных заявок по счёту.
  ///
  /// - Parameters:
  ///   - request: Request to send to GetOrders.
  ///   - callOptions: Call options.
  /// - Returns: A `UnaryCall` with futures for the metadata, status and response.
  public func getOrders(
    _ request: Tinkoff_Public_Invest_Api_Contract_V1_GetOrdersRequest,
    callOptions: CallOptions? = nil
  ) -> UnaryCall<Tinkoff_Public_Invest_Api_Contract_V1_GetOrdersRequest, Tinkoff_Public_Invest_Api_Contract_V1_GetOrdersResponse> {
    return self.makeUnaryCall(
      path: "/tinkoff.public.invest.api.contract.v1.OrdersService/GetOrders",
      request: request,
      callOptions: callOptions ?? self.defaultCallOptions,
      interceptors: self.interceptors?.makeGetOrdersInterceptors() ?? []
    )
  }
}

public protocol Tinkoff_Public_Invest_Api_Contract_V1_OrdersServiceClientInterceptorFactoryProtocol {

  /// - Returns: Interceptors to use when invoking 'postOrder'.
  func makePostOrderInterceptors() -> [ClientInterceptor<Tinkoff_Public_Invest_Api_Contract_V1_PostOrderRequest, Tinkoff_Public_Invest_Api_Contract_V1_PostOrderResponse>]

  /// - Returns: Interceptors to use when invoking 'cancelOrder'.
  func makeCancelOrderInterceptors() -> [ClientInterceptor<Tinkoff_Public_Invest_Api_Contract_V1_CancelOrderRequest, Tinkoff_Public_Invest_Api_Contract_V1_CancelOrderResponse>]

  /// - Returns: Interceptors to use when invoking 'getOrderState'.
  func makeGetOrderStateInterceptors() -> [ClientInterceptor<Tinkoff_Public_Invest_Api_Contract_V1_GetOrderStateRequest, Tinkoff_Public_Invest_Api_Contract_V1_OrderState>]

  /// - Returns: Interceptors to use when invoking 'getOrders'.
  func makeGetOrdersInterceptors() -> [ClientInterceptor<Tinkoff_Public_Invest_Api_Contract_V1_GetOrdersRequest, Tinkoff_Public_Invest_Api_Contract_V1_GetOrdersResponse>]
}

public final class Tinkoff_Public_Invest_Api_Contract_V1_OrdersServiceClient: Tinkoff_Public_Invest_Api_Contract_V1_OrdersServiceClientProtocol {
  public let channel: GRPCChannel
  public var defaultCallOptions: CallOptions
  public var interceptors: Tinkoff_Public_Invest_Api_Contract_V1_OrdersServiceClientInterceptorFactoryProtocol?

  /// Creates a client for the tinkoff.public.invest.api.contract.v1.OrdersService service.
  ///
  /// - Parameters:
  ///   - channel: `GRPCChannel` to the service host.
  ///   - defaultCallOptions: Options to use for each service call if the user doesn't provide them.
  ///   - interceptors: A factory providing interceptors for each RPC.
  public init(
    channel: GRPCChannel,
    defaultCallOptions: CallOptions = CallOptions(),
    interceptors: Tinkoff_Public_Invest_Api_Contract_V1_OrdersServiceClientInterceptorFactoryProtocol? = nil
  ) {
    self.channel = channel
    self.defaultCallOptions = defaultCallOptions
    self.interceptors = interceptors
  }
}

/// To build a server, implement a class that conforms to this protocol.
public protocol Tinkoff_Public_Invest_Api_Contract_V1_OrdersStreamServiceProvider: CallHandlerProvider {
  var interceptors: Tinkoff_Public_Invest_Api_Contract_V1_OrdersStreamServiceServerInterceptorFactoryProtocol? { get }

  ///Stream сделок пользователя
  func tradesStream(request: Tinkoff_Public_Invest_Api_Contract_V1_TradesStreamRequest, context: StreamingResponseCallContext<Tinkoff_Public_Invest_Api_Contract_V1_TradesStreamResponse>) -> EventLoopFuture<GRPCStatus>
}

extension Tinkoff_Public_Invest_Api_Contract_V1_OrdersStreamServiceProvider {
  public var serviceName: Substring { return "tinkoff.public.invest.api.contract.v1.OrdersStreamService" }

  /// Determines, calls and returns the appropriate request handler, depending on the request's method.
  /// Returns nil for methods not handled by this service.
  public func handle(
    method name: Substring,
    context: CallHandlerContext
  ) -> GRPCServerHandlerProtocol? {
    switch name {
    case "TradesStream":
      return ServerStreamingServerHandler(
        context: context,
        requestDeserializer: ProtobufDeserializer<Tinkoff_Public_Invest_Api_Contract_V1_TradesStreamRequest>(),
        responseSerializer: ProtobufSerializer<Tinkoff_Public_Invest_Api_Contract_V1_TradesStreamResponse>(),
        interceptors: self.interceptors?.makeTradesStreamInterceptors() ?? [],
        userFunction: self.tradesStream(request:context:)
      )

    default:
      return nil
    }
  }
}

public protocol Tinkoff_Public_Invest_Api_Contract_V1_OrdersStreamServiceServerInterceptorFactoryProtocol {

  /// - Returns: Interceptors to use when handling 'tradesStream'.
  ///   Defaults to calling `self.makeInterceptors()`.
  func makeTradesStreamInterceptors() -> [ServerInterceptor<Tinkoff_Public_Invest_Api_Contract_V1_TradesStreamRequest, Tinkoff_Public_Invest_Api_Contract_V1_TradesStreamResponse>]
}
/// Сервис предназначен для работы с торговыми поручениями:</br> **1**.
///выставление;</br> **2**. отмена;</br> **3**. получение статуса;</br> **4**.
///расчёт полной стоимости;</br> **5**. получение списка заявок.
///
/// To build a server, implement a class that conforms to this protocol.
public protocol Tinkoff_Public_Invest_Api_Contract_V1_OrdersServiceProvider: CallHandlerProvider {
  var interceptors: Tinkoff_Public_Invest_Api_Contract_V1_OrdersServiceServerInterceptorFactoryProtocol? { get }

  ///Метод выставления заявки.
  func postOrder(request: Tinkoff_Public_Invest_Api_Contract_V1_PostOrderRequest, context: StatusOnlyCallContext) -> EventLoopFuture<Tinkoff_Public_Invest_Api_Contract_V1_PostOrderResponse>

  ///Метод отмены биржевой заявки.
  func cancelOrder(request: Tinkoff_Public_Invest_Api_Contract_V1_CancelOrderRequest, context: StatusOnlyCallContext) -> EventLoopFuture<Tinkoff_Public_Invest_Api_Contract_V1_CancelOrderResponse>

  ///Метод получения статуса торгового поручения.
  func getOrderState(request: Tinkoff_Public_Invest_Api_Contract_V1_GetOrderStateRequest, context: StatusOnlyCallContext) -> EventLoopFuture<Tinkoff_Public_Invest_Api_Contract_V1_OrderState>

  ///Метод получения списка активных заявок по счёту.
  func getOrders(request: Tinkoff_Public_Invest_Api_Contract_V1_GetOrdersRequest, context: StatusOnlyCallContext) -> EventLoopFuture<Tinkoff_Public_Invest_Api_Contract_V1_GetOrdersResponse>
}

extension Tinkoff_Public_Invest_Api_Contract_V1_OrdersServiceProvider {
  public var serviceName: Substring { return "tinkoff.public.invest.api.contract.v1.OrdersService" }

  /// Determines, calls and returns the appropriate request handler, depending on the request's method.
  /// Returns nil for methods not handled by this service.
  public func handle(
    method name: Substring,
    context: CallHandlerContext
  ) -> GRPCServerHandlerProtocol? {
    switch name {
    case "PostOrder":
      return UnaryServerHandler(
        context: context,
        requestDeserializer: ProtobufDeserializer<Tinkoff_Public_Invest_Api_Contract_V1_PostOrderRequest>(),
        responseSerializer: ProtobufSerializer<Tinkoff_Public_Invest_Api_Contract_V1_PostOrderResponse>(),
        interceptors: self.interceptors?.makePostOrderInterceptors() ?? [],
        userFunction: self.postOrder(request:context:)
      )

    case "CancelOrder":
      return UnaryServerHandler(
        context: context,
        requestDeserializer: ProtobufDeserializer<Tinkoff_Public_Invest_Api_Contract_V1_CancelOrderRequest>(),
        responseSerializer: ProtobufSerializer<Tinkoff_Public_Invest_Api_Contract_V1_CancelOrderResponse>(),
        interceptors: self.interceptors?.makeCancelOrderInterceptors() ?? [],
        userFunction: self.cancelOrder(request:context:)
      )

    case "GetOrderState":
      return UnaryServerHandler(
        context: context,
        requestDeserializer: ProtobufDeserializer<Tinkoff_Public_Invest_Api_Contract_V1_GetOrderStateRequest>(),
        responseSerializer: ProtobufSerializer<Tinkoff_Public_Invest_Api_Contract_V1_OrderState>(),
        interceptors: self.interceptors?.makeGetOrderStateInterceptors() ?? [],
        userFunction: self.getOrderState(request:context:)
      )

    case "GetOrders":
      return UnaryServerHandler(
        context: context,
        requestDeserializer: ProtobufDeserializer<Tinkoff_Public_Invest_Api_Contract_V1_GetOrdersRequest>(),
        responseSerializer: ProtobufSerializer<Tinkoff_Public_Invest_Api_Contract_V1_GetOrdersResponse>(),
        interceptors: self.interceptors?.makeGetOrdersInterceptors() ?? [],
        userFunction: self.getOrders(request:context:)
      )

    default:
      return nil
    }
  }
}

public protocol Tinkoff_Public_Invest_Api_Contract_V1_OrdersServiceServerInterceptorFactoryProtocol {

  /// - Returns: Interceptors to use when handling 'postOrder'.
  ///   Defaults to calling `self.makeInterceptors()`.
  func makePostOrderInterceptors() -> [ServerInterceptor<Tinkoff_Public_Invest_Api_Contract_V1_PostOrderRequest, Tinkoff_Public_Invest_Api_Contract_V1_PostOrderResponse>]

  /// - Returns: Interceptors to use when handling 'cancelOrder'.
  ///   Defaults to calling `self.makeInterceptors()`.
  func makeCancelOrderInterceptors() -> [ServerInterceptor<Tinkoff_Public_Invest_Api_Contract_V1_CancelOrderRequest, Tinkoff_Public_Invest_Api_Contract_V1_CancelOrderResponse>]

  /// - Returns: Interceptors to use when handling 'getOrderState'.
  ///   Defaults to calling `self.makeInterceptors()`.
  func makeGetOrderStateInterceptors() -> [ServerInterceptor<Tinkoff_Public_Invest_Api_Contract_V1_GetOrderStateRequest, Tinkoff_Public_Invest_Api_Contract_V1_OrderState>]

  /// - Returns: Interceptors to use when handling 'getOrders'.
  ///   Defaults to calling `self.makeInterceptors()`.
  func makeGetOrdersInterceptors() -> [ServerInterceptor<Tinkoff_Public_Invest_Api_Contract_V1_GetOrdersRequest, Tinkoff_Public_Invest_Api_Contract_V1_GetOrdersResponse>]
}