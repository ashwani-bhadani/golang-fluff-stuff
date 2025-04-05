# golang-fluff-stuff
my itty bitty golang tutorial repo

## updates in GO since v1.17, currently v1.24.2
Here is a concise, point-wise summary of notable changes in Go from version 1.18 to 1.24.1:

**Go 1.18:**
- **Generics Introduced:** Added support for type parameters, enabling the creation of functions and data structures that operate on various types.
- **Fuzz Testing:** Integrated fuzzing into the standard testing framework for automated input generation to uncover edge cases.
- **Performance Enhancements:** Implemented register-based calling conventions on ARM64 and PowerPC64 architectures, improving performance.
- **Security Updates:** Disabled TLS 1.0 and 1.1 by default for clients; SHA-1 certificates are now rejected due to security concerns.

**Go 1.19:**
- **Soft Memory Limit:** Introduced a soft memory limit feature, allowing developers to set memory usage boundaries for applications.
- **Documentation Improvements:** Enhanced support for links, lists, and clearer headings in doc comments.
- **Runtime Optimizations:** Improved scheduling of garbage collection workers on idle threads and adjusted initial goroutine stack sizes based on historical data.

**Go 1.20:**
- **Profile-Guided Optimization (PGO):** Introduced PGO support, enabling performance tuning based on application profiles.
- **Code Coverage for Programs:** Extended code coverage tools to support whole-program analysis, not just unit tests.
- **Platform Support Changes:** Dropped support for older versions of Windows and macOS; added experimental support for FreeBSD on RISC-V.

**Go 1.21:**
- **New Standard Library Packages:** Added `slices` and `maps` packages for common operations on slices and maps.
- **Structured Logging:** Introduced the `log/slog` package for structured logging capabilities.
- **Language Enhancements:** Added support for type inference in `make` and `new` functions.

**Go 1.22:**
- **Runtime Improvements:** Enhanced garbage collection with more efficient metadata handling, leading to reduced memory overhead and improved CPU performance.
- **Compiler Optimizations:** Improved profile-guided optimization, resulting in significant runtime performance gains for many programs.

**Go 1.23:**
- **Security Enhancements:** Implemented stricter validation of HTTP headers to mitigate potential security vulnerabilities.
- **Performance Tweaks:** Optimized the `sync` package, reducing contention in high-concurrency scenarios.

**Go 1.24:**
- **Weak Pointers and Finalization:** Introduced the `weak` package for weak pointers and `runtime.AddCleanup` for improved finalization mechanisms.
- **Cryptographic Updates:** Added new packages `crypto/hkdf`, `crypto/pbkdf2`, and `crypto/sha3` for standardized cryptographic functions.
- **FIPS 140-3 Compliance:** Included mechanisms to facilitate compliance with FIPS 140-3 standards.
