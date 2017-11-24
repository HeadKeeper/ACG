#version 330

in vec3 vertexIn;
in vec2 vertTexCoord;
out vec2 fragTexCoord;

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

void main() {
    fragTexCoord = vertTexCoord;

    gl_Position = projection * camera * model * vec4(vertexIn, 1.0f);
}
