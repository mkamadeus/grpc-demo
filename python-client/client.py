import asyncio
import logging
import sys

import grpc
import student_pb2
import student_pb2_grpc


async def run(nim) -> None:
    async with grpc.aio.insecure_channel('localhost:1337') as channel:
        stub = student_pb2_grpc.StudentStub(channel)
        response = await stub.GetStudentByNIM(
            student_pb2.StudentByNIMRequest(nim=nim)
        )
    print("Client received:")
    print(str(response))


if __name__ == '__main__':
    if len(sys.argv) < 2:
        print("Usage: python client.py <NIM>")
        exit(-1)

    logging.basicConfig()
    asyncio.run(run(sys.argv[1]))