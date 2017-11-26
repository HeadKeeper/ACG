cmake_minimum_required(VERSION 3.8)
project(lab6)

set(CMAKE_CXX_STANDARD 11)

# sudo apt install libglfw3 libglfw3-dev libglew-dev libglm-dev libsoil-dev

set (
    SOURCE_FILES
        main.cpp
        Shader.h
        Camera.h
        ./libs/stb_image.h
)
add_executable(lab6 ${SOURCE_FILES})

configure_file(./shaders/lamp.vertex.shader ./shaders/lamp.vertex.shader COPYONLY)
configure_file(./shaders/lamp.fragment.shader ./shaders/lamp.fragment.shader COPYONLY)
configure_file(./shaders/lighting.vertex.shader ./shaders/lighting.vertex.shader COPYONLY)
configure_file(./shaders/lighting.fragment.shader ./shaders/lighting.fragment.shader COPYONLY)
configure_file(./textures/square.png ./textures/square.png COPYONLY)

# opengl
find_package(OpenGL REQUIRED)
include_directories(${OPENGL_INCLUDE_DIRS})

# glfw
find_package(glfw3 REQUIRED)
include_directories(${GLFW_INCLUDE_DIRS})
link_libraries(${GLFW_LIBRARY_DIRS})

# glew
find_package(GLEW REQUIRED)
include_directories(${GLEW_INCLUDE_DIRS})

target_link_libraries(lab6 ${OPENGL_LIBRARIES} ${GLEW_LIBRARIES} glfw)