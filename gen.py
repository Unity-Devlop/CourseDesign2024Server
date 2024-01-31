import argparse
import dataclasses
import os
import sys


@dataclasses.dataclass
class ProtoInfo:
    protoc_path: str
    csharp_grpc: str
    golang_grpc: str
    golang_protoc: str

    def __init__(self):
        pass


def gen_csharp(protoc_path, grpc_path, proto_dir, code_gen_dir):
    def gen(dir, gen_dir):
        # 移除之前生成的代码
        if os.path.exists(gen_dir):
            for root, dirs, files in os.walk(gen_dir):
                for file in files:
                    if file.endswith(".cs"):
                        os.remove(os.path.join(root, file))

        if not os.path.exists(gen_dir):
            os.makedirs(gen_dir)

        cmd = protoc_path + " -I=" + dir + " --csharp_out=" + gen_dir + " --grpc_out=" + gen_dir + " --plugin=protoc-gen-grpc=" + grpc_path + " " + dir + "/*.proto"
        os.system(cmd)

    # 一键生成顶级目录下的proto
    gen(proto_dir, code_gen_dir)
    # 遍历proto文件夹 找到子文件夹中所有的proto文件 如果存在嵌套文件夹则递归遍历
    for root, dirs, files in os.walk(proto_dir):
        for dir in dirs:
            if dir in ignore_dirs:
                continue
            gen(os.path.join(root, dir), os.path.join(code_gen_dir, dir))


def gen_go(protoc_path, protoc_go, grpc, proto_dir, code_gen_dir):
    def gen(dir, gen_dir):
        if os.path.exists(gen_dir):
            for root, dirs, files in os.walk(gen_dir):
                for file in files:
                    if file.endswith(".go"):
                        os.remove(os.path.join(root, file))
        if not os.path.exists(gen_dir):
            os.makedirs(gen_dir)
        cmd = protoc_path + " --go_out=" + gen_dir + " --go-grpc_out=" + gen_dir + " -I=" + dir + " " + dir + "/*.proto" + " --plugin=protoc-gen-go=" + protoc_go + " --plugin=protoc-gen-go-grpc=" + grpc
        os.system(cmd)


    gen(proto_dir, code_gen_dir)
    for root, dirs, files in os.walk(proto_dir):
        for dir in dirs:
            if dir in ignore_dirs:
                continue
            gen(os.path.join(root, dir), os.path.join(code_gen_dir, dir))


def get_proto_path() -> ProtoInfo:
    info = ProtoInfo()

    if os.name == "nt":
        info.protoc_path = "tools/protoc.exe"
        # csharp
        info.csharp_grpc = "./tools/grpc_csharp_plugin.exe"

        # golang
        info.golang_protoc = "./tools/protoc-gen-go.exe"
        info.golang_grpc = "./tools/protoc-gen-go-grpc.exe"
    else:
        raise "not support os"
    info.protoc_path = os.path.abspath(info.protoc_path)
    info.csharp_grpc = os.path.abspath(info.csharp_grpc)
    info.golang_protoc = os.path.abspath(info.golang_protoc)
    info.golang_grpc = os.path.abspath(info.golang_grpc)

    return info


if __name__ == "__main__":  # python gen.py --proto_dir=xxx --lang=xxx --gen_dir=xxx

    ignore_dirs: set = {"bin", "windows_x64"}
    info = get_proto_path()

    # 解析命令行参数
    parser = argparse.ArgumentParser()
    parser.add_argument("--proto_dir", type=str, default=None)
    parser.add_argument("--lang", type=str, default="go")
    parser.add_argument("--gen_dir", type=str, default=None)
    args = parser.parse_args()

    if args.proto_dir is None:
        print("proto_dir is None")
        sys.exit(-1)
    if args.lang is None:
        print("lang is None")
        sys.exit(-1)
    if args.gen_dir is None:
        args.gen_dir = args.proto_dir

    args.gen_dir = os.path.abspath(args.gen_dir)
    args.proto_dir = os.path.abspath(args.proto_dir)

    if args.lang == "csharp":
        # csharp
        gen_csharp(info.protoc_path, info.csharp_grpc, args.proto_dir, args.gen_dir)
    if args.lang == "go":
        # golang
        gen_go(info.protoc_path, info.golang_protoc, info.golang_grpc, args.proto_dir, args.gen_dir)
