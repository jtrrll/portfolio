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
        /// The database url to connect to.
        #[arg(long)]
        db_url: String,
        /// The path to save the resume.
        #[arg(long)]
        path: std::path::PathBuf,
    },
    /// Runs the web server.
    Run {
        /// The database url to connect to.
        #[arg(long)]
        db_url: String,
        /// The port number on which the server should listen.
        #[arg(long)]
        port: u16,
    },
}

/// Provides the portfolio CLI.
fn main() {
    let args = Args::parse();
    match args.command {
        Command::GenerateResume { path, .. } => {
            if !path.try_exists().unwrap_or(false) {
                Args::command()
                    .error(
                        clap::error::ErrorKind::ValueValidation,
                        &format!("The specified path '{}' does not exist.", path.display()),
                    )
                    .exit()
            };
            // TODO: Connect to database
            // TODO: Call the xp service with the db connection, and save the result
            println!("Done!");
        }
        Command::Run { port, .. } => {
            // TODO: Connect to database
            println!("Starting web server on port {}...", port);
            // TODO: Start the web server
        }
    }
}
