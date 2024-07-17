import type { ReactNode } from "react";
import { createContext, useContext, useEffect, useState } from "react";

type Theme = "dark" | "light";

interface ThemeProviderProps {
  children: ReactNode;
  defaultTheme?: Theme;
  storageKey?: string;
}

interface ThemeProps {
  theme: Theme;
  setTheme: (newTheme: Theme) => void;
}

const ThemeContext = createContext<ThemeProps>({
  theme: "light",
  setTheme: () => null,
});

export const ThemeProvider = ({
  children,
  storageKey = "ui-theme",
  ...props
}: ThemeProviderProps) => {
  const [theme, setTheme] = useState<Theme>(
    localStorage.getItem(storageKey) as Theme,
  );

  useEffect(() => {
    const root = window.document.documentElement;
    root.classList.remove("dark", "light");
    root.classList.add(localStorage.getItem(storageKey) as Theme);
  }, [theme]);

  const value = {
    theme,
    setTheme: (newTheme: Theme) => {
      setTheme(newTheme);
      localStorage.setItem(storageKey, newTheme);
    },
  };
  return (
    <ThemeContext.Provider value={value} {...props}>
      {children}
    </ThemeContext.Provider>
  );
};

export const useTheme = () => {
  const context = useContext(ThemeContext);
  return context;
};
