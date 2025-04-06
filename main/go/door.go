components {
  id: "door"
  component: "/main/scripts/door.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"door\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/game.atlas\"\n"
  "}\n"
  ""
  position {
    y: 40.0
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"default\"\n"
  "mask: \"default\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: 32.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 32.0\n"
  "  data: 32.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "label"
  type: "label"
  data: "size {\n"
  "  x: 180.0\n"
  "  y: 32.0\n"
  "}\n"
  "text: \"[E] to interact\"\n"
  "font: \"/main/fonts/acme.font\"\n"
  "material: \"/builtins/fonts/label.material\"\n"
  ""
  position {
    y: 100.0
  }
  scale {
    x: 0.7
    y: 0.7
  }
}
