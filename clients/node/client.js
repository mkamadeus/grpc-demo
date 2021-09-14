const grpc = require('@grpc/grpc-js')
const protoLoader = require('@grpc/proto-loader')

const main = () => {
  const packageDefinition = protoLoader.loadSync('../../schemas/student.proto')
  const StudentProto = grpc.loadPackageDefinition(packageDefinition).schemas
  const client = new StudentProto.Student('localhost:1337', grpc.credentials.createInsecure())
  
  client.GetStudentByNIM({nim: '13518035'}, (err, student) => {
    if(err) {
      console.log('Error:', err)
      return
    }

    console.log(student);
  })
}

main()