import os


def first_to_lower(s):
    s = chr(ord(s[0]) - ord('A') + ord('a')) + s[1:]
    return s


with open("generating/methodTemplate.txt", "r") as file:
    method_template = "".join(file.readlines())

with open("generating/protoTemplate.txt", "r") as file:
    proto_template = "".join(file.readlines())

with open("generating/serviceTemplate.txt", "r") as file:
    service_template = "".join(file.readlines())

functions = ["CreateGroupingFeed", "GetGroupingFeed", "GetGroupingFeedActivities", "UpdateGroupingFeed", "DeleteGroupingFeed"]

for name in functions:
    with open(f"internal/feeds/{first_to_lower(name)}.go", "w") as file:
        s = method_template.replace("MethodName", name)
        file.write(s)

with open("generating/proto.txt", "w") as file:
    for name in functions:
        file.write(proto_template.replace("MethodName", name))
    for name in functions:
        file.write(service_template.replace("MethodName", name))
