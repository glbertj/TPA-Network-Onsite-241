import type { Dispatch, ReactNode, SetStateAction } from "react";

import { Header } from "./Header.tsx";

export const Main = ({
  children,
  setSearch,
  search,
  setIsLoad,
}: {
  children: ReactNode;
  setSearch: Dispatch<SetStateAction<string>> | null;
  search?: string;
  setIsLoad?: Dispatch<SetStateAction<boolean>>;
}) => {
  return (
    <main id={"main"}>
      <Header setSearch={setSearch} search={search} setIsLoad={setIsLoad} />

      <div className={"content"} id={"content"}>
        {children}
      </div>
    </main>
  );
};
