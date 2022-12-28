use clap::{Parser, Subcommand};
mod workload;

#[derive(Parser, Debug)]
#[clap(author="Igor Miranda", version, about="Workload client")]
struct Arguments {
    #[clap(default_value_t=String::from("http://[::1]:8080"),short, long)]
    /// Server address
    address: String,
    #[clap(subcommand)]
    cmd: SubCommand,
}

#[derive(Subcommand, Debug)]
enum SubCommand {
    /// Request workload
    Request {
        #[clap(short, long)]
        /// Number of request
        count: u32
    },
    /// CPU workload
    Cpu {
        #[clap(short, long)]
        /// CPU usage percetage
        usage: u32,
        #[clap(short, long)]
        /// CPU usage duration in seconds
        duration: u64
    },
    /// Memory workload
    Memory {
        #[clap(short, long)]
        /// Quantity of memory to alloc in MB
        usage: u32,
        #[clap(short, long)]
        /// Memory usage time duration in seconds
        duration: u64,
    },
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let args = Arguments::parse();
    match args.cmd {
        SubCommand::Request { count } => {
            workload::request(args.address, count).await?;
        }
        SubCommand::Cpu { usage, duration } => {
            workload::cpu(args.address, usage, duration).await?;
        }
        SubCommand::Memory { usage, duration } => {
            workload::memory(args.address, usage, duration).await?;
        }
    }
    Ok(())
}

