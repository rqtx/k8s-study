mod hello;
mod workload;
use tonic::{transport::Server, transport::server::Router};

pub fn create_router() -> Router {
    return Server::builder()
        .add_service(hello::server())
        .add_service(workload::server());
}