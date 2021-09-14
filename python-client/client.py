import asyncio
import logging
import sys

import grpc
import student_pb2
import student_pb2_grpc


async def find_nim(nim) -> None:
    async with grpc.aio.insecure_channel('localhost:1337') as channel:
        stub = student_pb2_grpc.StudentStub(channel)
        response = await stub.GetStudentByNIM(
            student_pb2.StudentByNIMRequest(nim=nim)
        )
    print("Client received:")
    print(str(response))

async def find_faculty(fac) -> None:
    async with grpc.aio.insecure_channel('localhost:1337') as channel:
        stub = student_pb2_grpc.StudentStub(channel)
        request = student_pb2.StudentByFacultyRequest(faculty=fac)
        async for student in stub.GetStudentsByFaculty(request):
        	print(str(student))

if __name__ == '__main__':
    if len(sys.argv) < 2:
        print("Usage: python client.py <NIM>")
        exit(-1)

    logging.basicConfig()
    if len(sys.argv[1]) == 8 and sys.argv[1].isnumeric():
        asyncio.run(find_nim(sys.argv[1]))
    else:
        asyncio.run(find_faculty(sys.argv[1]))