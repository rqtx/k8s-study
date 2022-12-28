use std::error::Error;
use std::time::{Duration, SystemTime, UNIX_EPOCH};
use std::thread::sleep;
use tonic::{Request, Response, Status};
use workload::{WorkloadResponse, WorkloadRequest, workload_server::{Workload, WorkloadServer}};
use crate::config;

pub mod workload {
    tonic::include_proto!("workload");
}

#[derive(Debug, Default)]
pub struct WorkloadService {}

#[tonic::async_trait]
impl Workload for WorkloadService {
    async fn cpu(&self, request: Request<WorkloadRequest>) -> Result<Response<WorkloadResponse>, Status> {
        let req = request.into_inner();
        cpu(req.workload, req.duration);
        Ok(Response::new(workload::WorkloadResponse{}))
    }
    async fn memory(&self, request: Request<WorkloadRequest>) -> Result<Response<WorkloadResponse>, Status> {
        let req = request.into_inner();
        match memory(req.workload, req.duration) {
            Ok(_) => Ok(Response::new(workload::WorkloadResponse{})),
            Err(e) => Err(Status::new(tonic::Code::Internal, format!("{}", e))),
        }      
    }
}

pub fn server() -> WorkloadServer<WorkloadService>{
    WorkloadServer::new(WorkloadService::default())
}

fn cpu(workload: u32, duration: u64) {
    let now = || -> u128 {
        return SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_millis();
    };
    let duration_ms = Duration::from_secs(duration).as_millis();
    let time_of_run: f64 = 1000.0;
    let cpu_time_utilization: f64 = (workload as f64)/100.0;
    let on_time: u128 = (time_of_run * cpu_time_utilization) as u128;
    let off_time: u64 = (time_of_run * (1.0-cpu_time_utilization)) as u64;
    let start = now();
    let config = config::get_config();
    while now() - start < duration_ms {
        let work = now();
        while now() - work < on_time {
            fibonacci_reccursive(config.fibo);
        }
        sleep(Duration::from_millis(off_time));
    }
}

fn memory(workload: u32, duration: u64) -> Result<(), Box<dyn Error>> {
    let vecsize = ((workload as f64) * 1048576.0) as usize/std::mem::size_of::<i128>();
    let mut vec: Vec<i128> = Vec::new();
    vec.try_reserve_exact(vecsize)?;
    for i in 0..vecsize {
        vec.push(i as i128);
    }
    sleep(Duration::from_secs(duration));
    Ok(())
}

fn fibonacci_reccursive(n: u32) -> u64 {
	match n {
		0     => 0,
		1 | 2 => 1,
		3     => 2,
		_     => fibonacci_reccursive(n - 1) + fibonacci_reccursive(n - 2)
	}
}