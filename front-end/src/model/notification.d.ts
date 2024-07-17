interface NotificationSetting {
  notificationSettingId: string;
  emailAlbum: boolean;
  emailFollower: boolean;
  webAlbum: boolean;
  webFollower: boolean;
}

interface Notification {
  notifyId: string;
  userId: string;
  title: string;
  body: string;
  readAt: Date;
  status: string;
}
