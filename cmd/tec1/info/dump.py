#!/usr/bin/python3


import ezdxf


def extract_standalone_points(dxf_path):
    # Load the DXF file
    doc = ezdxf.readfile(dxf_path)
    msp = doc.modelspace()

    points = []
    # Query all POINT entities in the modelspace
    for entity in msp.query("POINT"):
        # Get the (x, y, z) tuple from the location attribute
        x, y, z = entity.dxf.location
        points.append((x, y, z))

    return points


def main():
    x_scale = 1.0 / 90.0
    y_scale = 1.0 / 122.0
    points = extract_standalone_points("./8seg.dxf")
    for i, p in enumerate(points):
        print(f"var p{i} = [2]float64{{ {p[0] * x_scale}, {p[1] * y_scale} }}")


main()
