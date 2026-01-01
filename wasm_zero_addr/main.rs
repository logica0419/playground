#![no_main]
#[unsafe(no_mangle)]
pub extern "C" fn test_zero() -> i32 {
  let ptr = 0 as *mut i32;

  unsafe {
    *ptr = 2025;
    return *ptr;
  }
}
