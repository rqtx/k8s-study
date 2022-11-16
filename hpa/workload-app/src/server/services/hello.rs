use tonic::{Request, Response, Status};
use hello::{HelloResponse, HelloRequest, hello_server::{Hello, HelloServer}};

pub mod hello {
    tonic::include_proto!("hello");
}

#[derive(Debug, Default)]
pub struct HelloService {}

#[tonic::async_trait]
impl Hello for HelloService {
    async fn say(&self, request: Request<HelloRequest>) -> Result<Response<HelloResponse>, Status> {
        let req = request.into_inner();
        Ok(Response::new(hello::HelloResponse {message: {format!("Hello {}", req.name)}}))
    }
}

pub fn server() -> HelloServer<HelloService>{
    return HelloServer::new(HelloService::default())
}