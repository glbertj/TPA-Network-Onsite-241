import { useContext } from "react";

import { SongContext } from "./SongProvider.tsx";

export const useSong = () => useContext(SongContext);
