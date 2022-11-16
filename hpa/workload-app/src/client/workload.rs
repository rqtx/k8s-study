use workload::{WorkloadRequest, workload_client::WorkloadClient};
use hello::{HelloRequest, hello_client::HelloClient};

pub mod workload {
    tonic::include_proto!("workload");
}

mod hello {
    tonic::include_proto!("hello");
}

pub async fn request(address: String, count: u32) -> Result<(), Box<dyn std::error::Error>> {
    let mut client = HelloClient::connect(address).await?;
    for _ in 1..=count {
        let request = tonic::Request::new(HelloRequest {name: String::from("World")});
        let response = client.say(request).await?;
        println!("{:?}", response.into_inner().message);
    }
    Ok(())
}

pub async fn cpu(address: String, usage: u32, duration: u64) -> Result<(), Box<dyn std::error::Error>> {
    let mut client = WorkloadClient::connect(address).await?;
    let request = tonic::Request::new(WorkloadRequest {workload: usage, duration: duration});
    client.cpu(request).await?;
    Ok(())
}

pub async fn memory(address: String, usage: u32, duration: u64) -> Result<(), Box<dyn std::error::Error>> {
    let mut client = WorkloadClient::connect(address).await?;
    let request = tonic::Request::new(WorkloadRequest {workload: usage, duration: duration});
    client.memory(request).await?;
    Ok(())
}