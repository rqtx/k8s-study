
mod services;
use services::{create_router};
mod config;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    
    let config = config::get_config();
    let socket = format!("{}:{}", config.addr, config.port).parse().unwrap();
    println!("Listening on {}", socket);
    create_router().serve(socket).await?;
    Ok(())
}

