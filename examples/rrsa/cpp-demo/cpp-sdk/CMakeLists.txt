cmake_minimum_required(VERSION 3.0)
project(cpp-example)
set(CMAKE_CXX_STANDARD 11)

add_executable(${PROJECT_NAME} src/demo.cc)
target_link_libraries(${PROJECT_NAME} alibabacloud-sdk-core)

target_link_libraries(${PROJECT_NAME})
target_include_directories(${PROJECT_NAME} PRIVATE /usr/include/jsoncpp)
target_link_libraries(${PROJECT_NAME} jsoncpp)

set_target_properties(${PROJECT_NAME}
  PROPERTIES
  OUTPUT_NAME ${TARGET_OUTPUT_NAME_PREFIX}${PROJECT_NAME}
  RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/bin)
