//! jtrrll's personal portfolio.

// The below lints prioritize increased strictness to raise code quality.
#![deny(
    anonymous_parameters,
    array_into_iter,
    bad_asm_style,
    bare_trait_objects,
    boxed_slice_into_iter,
    break_with_label_and_loop,
    deprecated,
    invalid_value,
    non_ascii_idents,
    non_fmt_panics,
    noop_method_call,
    unit_bindings,
    unused
)]
#![forbid(
    clippy::all,
    forbidden_lint_groups,
    keyword_idents,
    let_underscore,
    nonstandard_style,
    unsafe_code
)]
#![warn(
    clippy::cargo,
    clippy::pedantic,
    deprecated_safe,
    missing_copy_implementations,
    missing_debug_implementations,
    missing_docs,
    non_local_definitions,
    trivial_casts,
    trivial_numeric_casts,
    unused_results
)]

// Export library modules.
pub mod api;
pub mod html;
pub mod middleware;
mod services;
pub use services::*;
