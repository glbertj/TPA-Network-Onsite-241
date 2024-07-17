import { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";

export const useDebounce = <T>(value: T, delay = 1000): T => {
  const [debouncedValue, setDebouncedValue] = useState<T>(value);

  useEffect(() => {
    const handler = setTimeout(() => {
      // console.log(value);
      setDebouncedValue(value);
    }, delay);

    return () => {
      clearTimeout(handler);
    };
  }, [value, delay]);
  return debouncedValue;
};

export const useQuery = () => {
  return new URLSearchParams(useLocation().search);
};
