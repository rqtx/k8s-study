
mod services;
use services::{create_router};
use serde::Deserialize;

#[derive(Deserialize, Debug)]
struct Configuration {
    #[serde(default="default_port")]
    port: u16,
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let config = match envy::from_env::<Configuration>() {
        Ok(config) => config,
        Err(error) => panic!("{:#?}", error)
    };
    println!("Listening on {}:{}", "[::1]", config.port);
    let addr = format!("{}:{}", "[::1]", config.port).parse().unwrap();
    create_router().serve(addr).await?;
    Ok(())
}

fn default_port() -> u16 { 
    8080
} 