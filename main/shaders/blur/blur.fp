#version 140
in mediump vec2 var_texcoord0;
out vec4 out_fragColor;

uniform mediump sampler2D texture_sampler;

uniform fs_uniforms
{
    mediump vec4 tint;
    mediump vec4 blur_params;
    mediump vec4 texture_size;
};

void main()
{
    mediump vec4 tint_pm = vec4(tint.xyz * tint.w, tint.w);
    vec2 tex_size = textureSize(texture_sampler, 0);
    // vec2 tex_size = vec2(64.0);
    vec2 pixel_size = vec2(1.0 / tex_size.x, 1.0 / tex_size.y);
    
    out_fragColor = vec4(text_size.x, text_size.y, 1.0, 1.0);

    return;
    
    float blur_amount = blur_params.x;
    float blur_radius = max(0.0, blur_amount * 10.0);

    if (blur_radius < 0.01) {
        out_fragColor = texture(texture_sampler, var_texcoord0) * tint_pm;
        return;
    }

    vec4 color = vec4(0.0);
    float total_weight = 0.0;

    int samples = int(blur_radius * 2.0) + 1;
    samples = min(samples, 15);

    for (int x = -samples; x <= samples; x++) {
        for (int y = -samples; y <= samples; y++) {
            float dist = sqrt(float(x*x + y*y));
            float weight = exp(-(dist*dist) / (2.0 * blur_radius * blur_radius));
            vec2 offset = vec2(float(x), float(y)) * pixel_size;
            color += texture(texture_sampler, var_texcoord0 + offset) * weight;
            total_weight += weight;
        }
    }

    out_fragColor = vec4(pixel_size.x, pixel_size.y, 1.0, 1.0);
    // out_fragColor = (color / total_weight) * tint_pm;
}

