cmake_minimum_required(VERSION 2.8.12...3.29)

project(DemoCapiExtension)

# Core config
set(EXTENSION_NAME demo_capi)
set(LOADABLE_EXTENSION_NAME ${EXTENSION_NAME}_loadable_extension)
set(GO_EXTENSION_NAME go-extension)
set(BUILD_DIR ${CMAKE_BINARY_DIR}/extension/${EXTENSION_NAME}/${GO_EXTENSION_NAME}-prefix/src/${GO_EXTENSION_NAME}-build)
set(SOURCES add_numbers.cpp capi_demo.cpp)

include_directories(include)

# Configure extension: the minimum CAPI version where the extension still
# compiles should be chosen here
set(CAPI_MAJOR_VERSION 0)
set(CAPI_MINOR_VERSION 0)
set(CAPI_PATCH_VERSION 1)

set(CAPI_VERSION v${CAPI_MAJOR_VERSION}.${CAPI_MINOR_VERSION}.${CAPI_PATCH_VERSION})

# Set the version definitions. These prevent compilation if a function is used
# that is added in a version > CAPI_VERSION
add_definitions(-DDUCKDB_EXTENSION_API_VERSION_MAJOR=${CAPI_MAJOR_VERSION})
add_definitions(-DDUCKDB_EXTENSION_API_VERSION_MINOR=${CAPI_MINOR_VERSION})
add_definitions(-DDUCKDB_EXTENSION_API_VERSION_PATCH=${CAPI_PATCH_VERSION})
add_definitions(-DDUCKDB_EXTENSION_NAME=${EXTENSION_NAME})

include(ExternalProject)
ExternalProject_Add(
  ${GO_EXTENSION_NAME}
  SOURCE_DIR ${CMAKE_CURRENT_SOURCE_DIR}/go
  CONFIGURE_COMMAND ""
  UPDATE_COMMAND ""
  BUILD_COMMAND ${CMAKE_CURRENT_SOURCE_DIR}/go/build.sh "${CMAKE_CURRENT_SOURCE_DIR}/go/" "${CMAKE_SOURCE_DIR}/src/include/"
  BUILD_BYPRODUCTS "${BUILD_DIR}/${GO_EXTENSION_NAME}.a"
  INSTALL_COMMAND ""
  LOG_BUILD ON)

include_directories(${BUILD_DIR})

build_loadable_extension_capi(${EXTENSION_NAME} ${CAPI_VERSION} ${SOURCES})

target_link_libraries(
  ${LOADABLE_EXTENSION_NAME}
  debug
  "${BUILD_DIR}/${GO_EXTENSION_NAME}.a"
  optimized
  "${BUILD_DIR}/${GO_EXTENSION_NAME}.a"
  duckdb)

add_dependencies(${LOADABLE_EXTENSION_NAME} ${GO_EXTENSION_NAME})