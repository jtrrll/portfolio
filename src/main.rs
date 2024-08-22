use clap::{CommandFactory, Parser, Subcommand};

/// jtrrll's personal portfolio.
#[derive(Parser)]
#[command(version, about, long_about = None)]
struct Args {
    #[command(subcommand)]
    command: Command,
}

/// Commands available in the CLI.
#[derive(Subcommand)]
enum Command {
    /// Generates a resume as a PDF.
    GenerateResume {
        /// The path to save the resume.
        #[arg(short, long)]
        path: std::path::PathBuf,
    },
    /// Runs the web server.
    Run {
        /// The port number on which the server should listen.
        #[arg(short, long)]
        port: u16,
    },
}

/// Provides the portfolio CLI.
fn main() {
    let args = Args::parse();
    match args.command {
        Command::GenerateResume { path } => match path.exists() {
            true => {
                println!("TODO!! path {:?}", path)
            }
            false => Args::command()
                .error(
                    clap::error::ErrorKind::ValueValidation,
                    &format!("The specified path '{}' does not exist.", path.display()),
                )
                .exit(),
        },
        Command::Run { port } => {
            // TODO
            println!("port {}", port);
        }
    }
}
