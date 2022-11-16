fn main () -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::compile_protos("protos/hello.proto")?;
    tonic_build::compile_protos("protos/workload.proto")?;
    Ok(())
  }