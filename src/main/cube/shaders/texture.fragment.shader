#version 330

in vec2 fragTexCoord;
out vec4 outputColor;

uniform sampler2D textureUni;

void main()
{
    outputColor = texture(textureUni, fragTexCoord);
}
