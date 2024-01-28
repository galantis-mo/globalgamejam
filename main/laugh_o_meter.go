components {
  id: "laugh_o_meter"
  component: "/scripts/laugh_o_meter.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "outline"
  type: "sprite"
  data: "tile_set: \"/images/laugh_o_meter.atlas\"\n"
  "default_animation: \"laugh_o_meter_line\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ADD\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "inside"
  type: "sprite"
  data: "tile_set: \"/images/laugh_o_meter.atlas\"\n"
  "default_animation: \"laugh_o_meter_01\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_SCREEN\n"
  "size {\n"
  "  x: 575.0\n"
  "  y: 1794.0\n"
  "  z: 0.0\n"
  "  w: 0.0\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
