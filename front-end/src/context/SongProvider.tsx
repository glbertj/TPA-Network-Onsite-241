import type { AxiosResponse } from "axios";
import axios from "axios";
import { atom, useAtom } from "jotai";
import { atomWithStorage } from "jotai/utils";
import type { Dispatch, ReactNode, RefObject, SetStateAction } from "react";
import { useEffect } from "react";
import { useRef } from "react";
import { createContext, useState } from "react";

import { useAuth } from "./UseAuth.tsx";

export const SongContext = createContext<IProps>({} as IProps);

const queueAtom = atom<Song[] | null>(null);
const nowPlaying = atomWithStorage<Song | null>("nowPlaying", null);
const isUpdated = atom<boolean>(false);
const isPause = atomWithStorage<boolean>("paused", true);
const adv = atomWithStorage<number>("switch", 0);
const advertisement = atom<Advertisement | null>(null);

export const Enqueue = atom(null, (_, set, song: Song, user: User | null) => {
  if (user == null) return;
  void axios
    .post(
      "http://localhost:4000/auth/queue/enqueue?key=" + user.user_id,
      {
        songId: song.songId,
        title: song.title,
        artistId: song.artistId,
        albumId: song.albumId,
        releaseDate: song.releaseDate,
        duration: song.duration,
        file: song.file,
        play: song.play,
        artist: song.artist,
        album: song.album,
      },
      {
        withCredentials: true,
      },
    )
    .then(() => {
      set(isUpdated, true);
    });
});

export const RemoveQueue = atom(
  null,
  (_, set, index: number, user: User | null) => {
    if (user == null) return;
    void axios
      .post(
        "http://localhost:4000/auth/queue/remove?key=" +
          user.user_id +
          "&index=" +
          index.toString(),
        {},
        {
          withCredentials: true,
        },
      )
      .then(() => {
        set(isUpdated, true);
      });
  },
);

export const Dequeue = atom(null, (get, set, user: User | null) => {
  if (user == null) return;
  if (get(adv) < 5) {
    set(isPause, true);
    void axios
      .get("http://localhost:4000/auth/queue/dequeue?key=" + user.user_id, {
        withCredentials: true,
      })
      .then((res: AxiosResponse<WebResponse<Song>>) => {
        set(isUpdated, true);
        set(nowPlaying, res.data.data);
        set(isPause, false);
        set(adv, get(adv) + 1);
      })
      .catch((err: unknown) => {
        console.log(err);
        set(isPause, true);
      });
    return;
  }

  if (get(advertisement) == null) {
    void axios
      .get("http://localhost:4000/auth/adv/get", {
        withCredentials: true,
      })
      .then((res: AxiosResponse<WebResponse<Advertisement>>) => {
        set(advertisement, res.data.data);
        set(isPause, false);
        set(adv, get(adv) + 1);
      })
      .catch(() => {
        console.log("error");
      });
  }
});

export const ResetAdv = atom(null, (get, set) => {
  if (get(adv) >= 5 || get(advertisement) != null) {
    set(adv, 0);
    set(advertisement, null);
  }
});

export const ClearQueue = atom(null, (_, set, user: User | null) => {
  if (user == null) return;
  localStorage.setItem("nowPlaying", "");
  void axios
    .get("http://localhost:4000/auth/queue/clear?key=" + user.user_id, {
      withCredentials: true,
    })
    .then(() => {
      set(isUpdated, true);
      set(queueAtom, []);
    });
});

export const GetAllQueue = atom(null, (_, set, user: User | null) => {
  if (user == null) return;
  void axios
    .get("http://localhost:4000/auth/queue/get-all?key=" + user.user_id, {
      withCredentials: true,
    })
    .then((res: AxiosResponse<WebResponse<Song[]>>) => {
      set(isUpdated, false);
      set(queueAtom, res.data.data);
    })
    .catch((err: unknown) => {
      console.log(err);
    });
});

interface IProps {
  song: Song | null;
  changeSong: (songId: string) => void;
  showDetail: string;
  showDetailHandler: (type: string) => void;
  track: string;
  setSong: Dispatch<SetStateAction<Song | null>>;
  setTrack: Dispatch<SetStateAction<string>>;
  handlePlay: () => void;
  isPaused: boolean;
  playlist: Playlist[] | undefined;
  updatePlaylist: () => void;
  enqueue: (song: Song, user: User | null) => void;
  dequeue: (user: User | null) => void;
  clearAllQueue: () => void;
  getAllQueue: (user: User | null) => void;
  audioRef: RefObject<HTMLAudioElement | null>;
  waitingSong: Song[] | null;
  removeQueue: (index: number, user: User | null) => void;
  advertise: Advertisement | null;
  resetAdv: () => void;
  closeAdvertise: () => void;
}

export const SongProvider = ({ children }: { children: ReactNode }) => {
  const { user } = useAuth();
  const [song, setSong] = useAtom(nowPlaying);
  const [playlist, setPlaylist] = useState<Playlist[]>();
  const [, enqueue] = useAtom(Enqueue);
  const [, dequeue] = useAtom(Dequeue);
  const [, clearQueue] = useAtom(ClearQueue);
  const [, getAllQueue] = useAtom(GetAllQueue);
  const [, removeQueues] = useAtom(RemoveQueue);
  const [, resetAdv] = useAtom(ResetAdv);
  const audioRef = useRef<HTMLAudioElement | null>(null);
  const [waitingSong] = useAtom(queueAtom);
  const [update] = useAtom(isUpdated);
  const [isPaused, setIsPaused] = useAtom(isPause);
  const [advCount, setAdvCount] = useAtom(adv);
  const [advertise] = useAtom(advertisement);

  const removeQueue = (index: number, user: User | null) => {
    if (user == null) return;
    removeQueues(index, user);
  };

  useEffect(() => {
    if (user == null) return;
    if (advCount > 5) {
      dequeue(user);
      setShowDetail("advertise");
    }
  }, [advCount, setAdvCount, user]);

  useEffect(() => {
    if (song != null) return;
    dequeue(user);
  }, [user]);

  useEffect(() => {
    if (update) {
      if (song == null) {
        dequeue(user);
      }
      getAllQueue(user);
    }
  }, [update, user]);

  useEffect(() => {
    updatePlaylist();
    getAllQueue(user);
  }, [user]);

  useEffect(() => {
    if (song == null || advCount > 5) return;
    audioRef.current?.pause();
    audioRef.current = new Audio(
      "http://localhost:4000/auth/music?id=" + song.songId,
    );
    // audioRef.current.preload = "auto";
    if (!isPaused) {
      audioRef.current.play().catch((error: unknown) => {
        console.log(error);
        return;
      });
    }
    // axios
    //   .get("http://localhost:4000/test?id=" + song.songId, {
    //     responseType: "blob",
    //     headers: {
    //       "Content-Type": "audio/mpeg",
    //     },
    //   })
    //   .then((response: AxiosResponse<Blob>) => {
    //     const blob = response.data;
    //     const audioURL = URL.createObjectURL(blob);
    //     if (song.songAudio) {
    //       song.songAudio.src = audioURL;
    //     } else {
    //       song.songAudio = new Audio(audioURL);
    //     }
    //     audioRef.current?.pause();
    //     audioRef.current = null;
    //     audioRef.current = song.songAudio;
    //     if (!isPaused) {
    //       audioRef.current.play().catch((error: unknown) => {
    //         console.log(error);
    //         return;
    //       });
    //     }
    //   })
    //   .catch((error: unknown) => {
    //     console.error("Error fetching music:", error);
    //   });
  }, [setIsPaused, song, advCount]);

  useEffect(() => {
    console.log(advertise);
    if (advertise == null) return;
    audioRef.current?.pause();
    audioRef.current = new Audio(
      "http://localhost:4000/auth/adv?id=" + advertise.advertisementId,
    );
    audioRef.current.preload = "metadata";
    audioRef.current.play().catch((error: unknown) => {
      console.log(error);
      return;
    });
    // axios
    //   .get("http://localhost:4000/adv?id=" + advertise.advertisementId, {
    //     responseType: "blob",
    //   })
    //   .then((response: AxiosResponse<Blob>) => {
    //     const blob = response.data;
    //     const audioURL = URL.createObjectURL(blob);
    //     // if (song.songAudio) {
    //     //   song.songAudio.src = audioURL;
    //     // } else {
    //     //   song.songAudio = new Audio(audioURL);
    //     // }
    //     audioRef.current?.pause();
    //     audioRef.current = null;
    //     audioRef.current = new Audio(audioURL);
    //     // audioRef.current.muted = true;
    //     audioRef.current.play().catch((error: unknown) => {
    //       console.log(error);
    //       return;
    //     });
    //   })
    //   .catch((error: unknown) => {
    //     console.error("Error fetching music:", error);
    //   });
  }, [advertise]);

  const clearAllQueue = () => {
    clearQueue(user);
    setSong(null);
  };

  const updatePlaylist = () => {
    if (user == null) return;
    axios
      .get("http://localhost:4000/auth/playlist?id=" + user.user_id, {
        withCredentials: true,
      })
      .then((res: AxiosResponse<WebResponse<Playlist[]>>) => {
        setPlaylist(res.data.data);
        // console.log(res.data.data)
      })
      .catch((err: unknown) => {
        console.log(err);
      });
  };

  const handlePlay = () => {
    if (audioRef.current == null) return;
    if (advCount > 5) return;
    if (!isPaused) {
      // audioRef.current.play().catch((error: unknown) => {
      //   console.log(error);
      //   return;
      // });
      setIsPaused(true);
    } else {
      // audioRef.current.pause();
      setIsPaused(false);
    }
  };

  const [track, setTrack] = useState<string>("");

  const [showDetail, setShowDetail] = useState<string>("");

  const showDetailHandler = (type: string) => {
    if (advCount >= 5 && type !== "advertise") return;
    if (showDetail === type) {
      setShowDetail("");
    } else {
      setShowDetail(type);
    }
  };

  const closeAdvertise = () => {
    setShowDetail("");
  };

  const changeSong = (songId: string) => {
    if (user == null) return;
    axios
      .get("http://localhost:4000/auth/song/get?id=" + songId, {
        withCredentials: true,
      })
      .then((res: AxiosResponse<WebResponse<Song>>) => {
        // console.log(res.data);
        // setSong(res.data.data);
        clearQueue(user);
        enqueue(res.data.data, user);
      })
      .catch((err: unknown) => {
        console.log(err);
      });
  };

  const values: IProps = {
    setSong,
    song,
    changeSong,
    showDetail,
    showDetailHandler,
    track,
    setTrack,
    handlePlay,
    isPaused,
    playlist,
    updatePlaylist,
    enqueue,
    dequeue,
    clearAllQueue,
    getAllQueue,
    audioRef,
    waitingSong,
    removeQueue,
    advertise,
    resetAdv,
    closeAdvertise,
  };

  return <SongContext.Provider value={values}>{children}</SongContext.Provider>;
};
