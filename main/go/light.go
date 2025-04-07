components {
  id: "script"
  component: "/main/scripts/light.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"light\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/backgrounds.atlas\"\n"
  "}\n"
  ""
}
