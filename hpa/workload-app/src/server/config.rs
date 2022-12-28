use serde::Deserialize;

#[derive(Deserialize, Debug)]
pub struct Configuration {
    #[serde(default="default_port")]
    pub port: u16,
    #[serde(default="default_addr")]
    pub addr: String,
    #[serde(default="default_fibo")]
    pub fibo: u32
}

pub fn get_config() -> Configuration{
    match envy::from_env::<Configuration>() {
        Ok(config) => config,
        Err(error) => panic!("{:#?}", error)
    }
}

fn default_addr() -> String { 
    String::from("[::1]")
}

fn default_port() -> u16 { 
    8080
} 

fn default_fibo() -> u32 { 
    25
} 