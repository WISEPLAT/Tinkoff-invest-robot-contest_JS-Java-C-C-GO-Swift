cmake_minimum_required(VERSION 3.11)

project(TinkoffInvestSDK)

include(GenerateExportHeader)
include(cmake/common.cmake)

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -fPIC")

get_filename_component(hw_proto "${CMAKE_CURRENT_SOURCE_DIR}/investAPI/src/docs/contracts/*.proto" ABSOLUTE)
get_filename_component(hw_proto_path "${hw_proto}" PATH)

file(GLOB tinkoff_proto_files "${hw_proto_path}/*.proto")

string(REPLACE ${hw_proto_path} ${CMAKE_CURRENT_BINARY_DIR} OUTPUT_FILE_NAMES "${tinkoff_proto_files}")

string(REPLACE ".proto" ".pb.cc" OUTPUT_PB_SOURCES "${OUTPUT_FILE_NAMES}")
string(REPLACE ".proto" ".grpc.pb.cc" OUTPUT_GRPC_SOURCES "${OUTPUT_FILE_NAMES}")
string(REPLACE ".proto" ".pb.h" OUTPUT_PB_HEADERS "${OUTPUT_FILE_NAMES}")
string(REPLACE ".proto" ".grpc.pb.h" OUTPUT_GRPC_HEADERS "${OUTPUT_FILE_NAMES}")

add_custom_command(
      OUTPUT ${OUTPUT_PB_SOURCES} ${OUTPUT_PB_HEADERS} ${OUTPUT_GRPC_SOURCES} ${OUTPUT_GRPC_HEADERS}
      COMMAND ${_PROTOBUF_PROTOC}
      ARGS --grpc_out "${CMAKE_CURRENT_BINARY_DIR}"
        --proto_path ${CMAKE_BINARY_DIR}/_deps/grpc-src/third_party/protobuf/src
        --cpp_out "${CMAKE_CURRENT_BINARY_DIR}"
        -I "${hw_proto_path}"
        --plugin=protoc-gen-grpc="${_GRPC_CPP_PLUGIN_EXECUTABLE}"
        "${hw_proto}"
      COMMAND sh ${CMAKE_CURRENT_SOURCE_DIR}/cmake/fix_keywords.sh ${CMAKE_CURRENT_BINARY_DIR}
      DEPENDS "${hw_proto}")

add_library(hw_grpc_proto
  ${OUTPUT_PB_SOURCES}
  ${OUTPUT_PB_HEADERS}
  ${OUTPUT_GRPC_SOURCES}
  ${OUTPUT_GRPC_HEADERS}
  )

target_link_libraries(hw_grpc_proto
  ${_REFLECTION}
  ${_GRPC_GRPCPP}
  ${_PROTOBUF_LIBPROTOBUF})

file(GLOB TARGET_SRC "${CMAKE_CURRENT_SOURCE_DIR}/*.cpp" "${CMAKE_CURRENT_SOURCE_DIR}/services/*.cpp")
file(GLOB TARGET_HDR "${CMAKE_CURRENT_SOURCE_DIR}/*.h" "${CMAKE_CURRENT_SOURCE_DIR}/services/*.h")

add_library(${PROJECT_NAME} SHARED ${TARGET_SRC} ${TARGET_HDR})

generate_export_header(${PROJECT_NAME})

target_include_directories(${PROJECT_NAME}
    PUBLIC
        ${CMAKE_CURRENT_BINARY_DIR}
        ${CMAKE_CURRENT_SOURCE_DIR}
        ${CMAKE_CURRENT_BINARY_DIR}/services
        ${CMAKE_CURRENT_SOURCE_DIR}/services
    )

target_link_libraries(${PROJECT_NAME}
    PRIVATE
    hw_grpc_proto
    ${_REFLECTION}
    ${_GRPC_GRPCPP}
    ${_PROTOBUF_LIBPROTOBUF}
    )

add_subdirectory(samples)
