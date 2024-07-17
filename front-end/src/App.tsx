import "./css/App.scss";

import { GoogleOAuthProvider } from "@react-oauth/google";
import { Provider } from "jotai";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";

import { AuthProvider } from "./context/AuthProvider.tsx";
import { SongProvider } from "./context/SongProvider.tsx";
// import {ChatPage} from "./route/ChatPage.tsx";
import { ThemeProvider } from "./context/ThemeProvider.tsx";
import { AlbumPage } from "./route/AlbumPage.tsx";
import { ArtistPage } from "./route/ArtistPage.tsx";
import { ArtistVerificationPage } from "./route/ArtistVerificationPage.tsx";
import { CreateMusicPage } from "./route/CreateMusicPage.tsx";
import { CreatePlaylistPage } from "./route/CreatePlaylistPage.tsx";
import { EditProfilePage } from "./route/EditProfilePage.tsx";
import { ForgotPasswordPage } from "./route/ForgotPasswordPage.tsx";
import { GetVerifiedPage } from "./route/GetVerifiedPage.tsx";
import { HomePage } from "./route/HomePage.tsx";
import { LoginPage } from "./route/LoginPage.tsx";
import { NotFoundPage } from "./route/NotFoundPage.tsx";
import { NotificationPage } from "./route/NotificationPage.tsx";
import { PlaylistPage } from "./route/PlaylistPage.tsx";
import { ProfilePage } from "./route/ProfilePage.tsx";
import { RegisterPage } from "./route/RegisterPage.tsx";
import { ResetPasswordPage } from "./route/ResetPasswordPage.tsx";
import { SearchPage } from "./route/SearchPage.tsx";
import { SettingPage } from "./route/SettingPage.tsx";
import { ShowMorePage } from "./route/ShowMorePage.tsx";
import { TrackPage } from "./route/TrackPage.tsx";
import { VerifyEmail } from "./route/VerifyEmail.tsx";
import { YourPostPage } from "./route/YourPostPage.tsx";

function App() {
  return (
    <Router>
      <GoogleOAuthProvider
        clientId={
          "841490014876-nc1omea8apsbevhcouqgbhq9mmc5tq7k.apps.googleusercontent.com"
        }
      >
        <ThemeProvider>
          <AuthProvider>
            <SongProvider>
              <Provider>
                <Routes>
                  <Route path="/home" element={<HomePage />} />
                  <Route path="/search" element={<SearchPage />} />
                  <Route path="/profile/:id" element={<ProfilePage />} />
                  <Route path="/artist/:id" element={<ArtistPage />} />
                  <Route path="/track/:id" element={<TrackPage />} />
                  <Route
                    path="/more/:type/:subtype/:id"
                    element={<ShowMorePage />}
                  />
                  <Route path="/album/:id" element={<AlbumPage />} />
                  <Route path="/playlist/:id" element={<PlaylistPage />} />
                  <Route path="/user/edit/" element={<EditProfilePage />} />
                  <Route path="/get-verified/" element={<GetVerifiedPage />} />
                  <Route path="/account/settings" element={<SettingPage />} />
                  <Route
                    path="/artist/verif"
                    element={<ArtistVerificationPage />}
                  />
                  <Route path="/your-post" element={<YourPostPage />} />
                  <Route path="/create/music" element={<CreateMusicPage />} />
                  <Route path="/login" element={<LoginPage />} />
                  <Route path="/register" element={<RegisterPage />} />
                  <Route path="/forgot" element={<ForgotPasswordPage />} />
                  <Route path="/reset-pass" element={<ResetPasswordPage />} />
                  <Route path="/verify-email" element={<VerifyEmail />} />
                  <Route
                    path="/notification/setting"
                    element={<NotificationPage />}
                  />
                  <Route
                    path={"/playlist/create"}
                    element={<CreatePlaylistPage />}
                  />
                  {/*<Route path="/chat" element={<ChatPage/>} />*/}
                  <Route path="/404" element={<NotFoundPage />} />
                  <Route path="*" element={<HomePage />} />
                </Routes>
              </Provider>
            </SongProvider>
          </AuthProvider>
        </ThemeProvider>
      </GoogleOAuthProvider>
    </Router>
  );
}

export default App;
