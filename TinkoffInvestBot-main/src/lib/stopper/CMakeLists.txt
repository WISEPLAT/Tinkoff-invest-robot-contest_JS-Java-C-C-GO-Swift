cmake_minimum_required(VERSION 3.11)

project(stopper)

set(CMAKE_CXX_STANDARD 14)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_AUTOMOC ON)

find_package(Qt5 REQUIRED COMPONENTS Core)

add_library(${PROJECT_NAME} SHARED
    stopper.h
    stopper.cpp
)

target_include_directories(${PROJECT_NAME}
    PUBLIC
        ${CMAKE_CURRENT_BINARY_DIR}
        ${CMAKE_CURRENT_SOURCE_DIR}
    )

target_link_libraries(${PROJECT_NAME}
    PRIVATE
    Qt5::Core
    )

add_subdirectory(tests)
