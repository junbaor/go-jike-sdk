package jike

import "time"

type LoginOutput struct {
	IsRegister bool `json:"isRegister"`
	User       struct {
		Id            string    `json:"id"`
		Username      string    `json:"username"`
		ScreenName    string    `json:"screenName"`
		CreatedAt     time.Time `json:"createdAt"`
		UpdatedAt     time.Time `json:"updatedAt"`
		IsVerified    bool      `json:"isVerified"`
		VerifyMessage string    `json:"verifyMessage"`
		BriefIntro    string    `json:"briefIntro"`
		AvatarImage   struct {
			ThumbnailUrl string `json:"thumbnailUrl"`
			SmallPicUrl  string `json:"smallPicUrl"`
			MiddlePicUrl string `json:"middlePicUrl"`
			PicUrl       string `json:"picUrl"`
			Format       string `json:"format"`
		} `json:"avatarImage"`
		Decorations struct {
		} `json:"decorations"`
		ProfileImageUrl string `json:"profileImageUrl"`
		StatsCount      struct {
			TopicSubscribed            int `json:"topicSubscribed"`
			TopicCreated               int `json:"topicCreated"`
			FollowedCount              int `json:"followedCount"`
			FollowingCount             int `json:"followingCount"`
			HighlightedPersonalUpdates int `json:"highlightedPersonalUpdates"`
			Liked                      int `json:"liked"`
			RespectedCount             int `json:"respectedCount"`
		} `json:"statsCount"`
		IsBannedForever bool `json:"isBannedForever"`
		IsSponsor       bool `json:"isSponsor"`
		BackgroundImage struct {
			PicUrl string `json:"picUrl"`
		} `json:"backgroundImage"`
		Bio         string `json:"bio"`
		Gender      string `json:"gender"`
		GroupId     string `json:"groupId"`
		Preferences struct {
			PersonalizedRecommendationOn  bool        `json:"personalizedRecommendationOn"`
			AutoPlayVideo                 bool        `json:"autoPlayVideo"`
			TopicTagGuideOn               bool        `json:"topicTagGuideOn"`
			DailyNotificationOn           bool        `json:"dailyNotificationOn"`
			FollowedNotificationOn        bool        `json:"followedNotificationOn"`
			ReplyNotificationAllowGroup   string      `json:"replyNotificationAllowGroup"`
			LikeNotificationOn            bool        `json:"likeNotificationOn"`
			RespectNotificationOn         bool        `json:"respectNotificationOn"`
			MentionNotificationOn         bool        `json:"mentionNotificationOn"`
			LiveNotificationOn            bool        `json:"liveNotificationOn"`
			SubscribeWeatherForecast      bool        `json:"subscribeWeatherForecast"`
			PrivateTopicSubscribe         bool        `json:"privateTopicSubscribe"`
			UndiscoverableByPhoneNumber   bool        `json:"undiscoverableByPhoneNumber"`
			SaveDataUsageMode             bool        `json:"saveDataUsageMode"`
			TopicPushSettings             string      `json:"topicPushSettings"`
			DebugLogOn                    bool        `json:"debugLogOn"`
			Env                           string      `json:"env"`
			SyncCommentToPersonalActivity bool        `json:"syncCommentToPersonalActivity"`
			RepostWithComment             bool        `json:"repostWithComment"`
			EnablePrivateChat             bool        `json:"enablePrivateChat"`
			BlockStrangerPoke             bool        `json:"blockStrangerPoke"`
			EnablePictureWatermark        bool        `json:"enablePictureWatermark"`
			EnableOperationPush           bool        `json:"enableOperationPush"`
			EnableChatSound               bool        `json:"enableChatSound"`
			OpenMessageTabOnLaunch        bool        `json:"openMessageTabOnLaunch"`
			TabOnLaunch                   string      `json:"tabOnLaunch"`
			HideSubscribeTab              interface{} `json:"hideSubscribeTab"`
			UndiscoverableByWeiboUser     interface{} `json:"undiscoverableByWeiboUser"`
			PaidMarket                    interface{} `json:"paidMarket"`
			FollowingUpdatesSortBy        string      `json:"followingUpdatesSortBy"`
		} `json:"preferences"`
		IsBetaUser         bool      `json:"isBetaUser"`
		UsernameSet        bool      `json:"usernameSet"`
		WechatOpenId       string    `json:"wechatOpenId"`
		WechatUnionId      string    `json:"wechatUnionId"`
		MobilePhoneNumber  string    `json:"mobilePhoneNumber"`
		AreaCode           string    `json:"areaCode"`
		GroupVersion       int       `json:"groupVersion"`
		LastChangeNameTime time.Time `json:"lastChangeNameTime"`
		Birthday           string    `json:"birthday"`
		WechatUserInfo     struct {
			Id           string `json:"id"`
			ExternalName string `json:"externalName"`
		} `json:"wechatUserInfo"`
		TopicRoles  []interface{} `json:"topicRoles"`
		Zodiac      string        `json:"zodiac"`
		ShowRespect bool          `json:"showRespect"`
		Medals      []interface{} `json:"medals"`
		ProfileTags []struct {
			PicUrl string `json:"picUrl,omitempty"`
			Type   string `json:"type"`
			Text   string `json:"text,omitempty"`
		} `json:"profileTags"`
		BackgroundMaskColor    string `json:"backgroundMaskColor"`
		IsLoginUser            bool   `json:"isLoginUser"`
		IsBanned               bool   `json:"isBanned"`
		UserHasPosted          bool   `json:"userHasPosted"`
		RegisterAppVersion     string `json:"registerAppVersion"`
		HasStories             bool   `json:"hasStories"`
		RestrictedAvatarChange struct {
			NextChangeDate string `json:"nextChangeDate"`
			Limits         int    `json:"limits"`
		} `json:"restrictedAvatarChange"`
		RestrictedNameChange struct {
			NextChangeDate string `json:"nextChangeDate"`
			Limits         int    `json:"limits"`
		} `json:"restrictedNameChange"`
		IsDefaultScreenName bool `json:"isDefaultScreenName"`
		ProfileVisitInfo    struct {
			TodayCount    int `json:"todayCount"`
			LatestVisitor struct {
				Id            string    `json:"id"`
				Username      string    `json:"username"`
				ScreenName    string    `json:"screenName"`
				CreatedAt     time.Time `json:"createdAt"`
				UpdatedAt     time.Time `json:"updatedAt"`
				IsVerified    bool      `json:"isVerified"`
				VerifyMessage string    `json:"verifyMessage"`
				BriefIntro    string    `json:"briefIntro"`
				AvatarImage   struct {
					ThumbnailUrl string `json:"thumbnailUrl"`
					SmallPicUrl  string `json:"smallPicUrl"`
					MiddlePicUrl string `json:"middlePicUrl"`
					PicUrl       string `json:"picUrl"`
					Format       string `json:"format"`
				} `json:"avatarImage"`
				Decorations struct {
				} `json:"decorations"`
				ProfileImageUrl string `json:"profileImageUrl"`
				StatsCount      struct {
					TopicSubscribed            int `json:"topicSubscribed"`
					TopicCreated               int `json:"topicCreated"`
					FollowedCount              int `json:"followedCount"`
					FollowingCount             int `json:"followingCount"`
					HighlightedPersonalUpdates int `json:"highlightedPersonalUpdates"`
					Liked                      int `json:"liked"`
					RespectedCount             int `json:"respectedCount"`
				} `json:"statsCount"`
				IsBannedForever bool   `json:"isBannedForever"`
				IsSponsor       bool   `json:"isSponsor"`
				Bio             string `json:"bio"`
				Gender          string `json:"gender"`
				City            string `json:"city"`
				Country         string `json:"country"`
				Province        string `json:"province"`
			} `json:"latestVisitor"`
		} `json:"profileVisitInfo"`
	} `json:"user"`
	EnabledFeatures []string `json:"enabledFeatures"`
	AgreedProtocol  []string `json:"agreedProtocol"`
}

type ProfileOutput struct {
	User struct {
		Id            string    `json:"id"`
		Username      string    `json:"username"`
		ScreenName    string    `json:"screenName"`
		CreatedAt     time.Time `json:"createdAt"`
		UpdatedAt     time.Time `json:"updatedAt"`
		IsVerified    bool      `json:"isVerified"`
		VerifyMessage string    `json:"verifyMessage"`
		BriefIntro    string    `json:"briefIntro"`
		AvatarImage   struct {
			ThumbnailUrl string `json:"thumbnailUrl"`
			SmallPicUrl  string `json:"smallPicUrl"`
			MiddlePicUrl string `json:"middlePicUrl"`
			PicUrl       string `json:"picUrl"`
			Format       string `json:"format"`
		} `json:"avatarImage"`
		Decorations struct {
		} `json:"decorations"`
		ProfileImageUrl string `json:"profileImageUrl"`
		StatsCount      struct {
			TopicSubscribed            int `json:"topicSubscribed"`
			TopicCreated               int `json:"topicCreated"`
			FollowedCount              int `json:"followedCount"`
			FollowingCount             int `json:"followingCount"`
			HighlightedPersonalUpdates int `json:"highlightedPersonalUpdates"`
			Liked                      int `json:"liked"`
			RespectedCount             int `json:"respectedCount"`
		} `json:"statsCount"`
		IsBannedForever bool `json:"isBannedForever"`
		IsSponsor       bool `json:"isSponsor"`
		BackgroundImage struct {
			PicUrl string `json:"picUrl"`
		} `json:"backgroundImage"`
		Bio         string `json:"bio"`
		Gender      string `json:"gender"`
		GroupId     string `json:"groupId"`
		Preferences struct {
			AutoPlayVideo                 bool        `json:"autoPlayVideo"`
			TopicTagGuideOn               bool        `json:"topicTagGuideOn"`
			DailyNotificationOn           bool        `json:"dailyNotificationOn"`
			FollowedNotificationOn        bool        `json:"followedNotificationOn"`
			SubscribeWeatherForecast      bool        `json:"subscribeWeatherForecast"`
			PrivateTopicSubscribe         bool        `json:"privateTopicSubscribe"`
			UndiscoverableByPhoneNumber   bool        `json:"undiscoverableByPhoneNumber"`
			SaveDataUsageMode             bool        `json:"saveDataUsageMode"`
			TopicPushSettings             string      `json:"topicPushSettings"`
			DebugLogOn                    bool        `json:"debugLogOn"`
			Env                           string      `json:"env"`
			SyncCommentToPersonalActivity bool        `json:"syncCommentToPersonalActivity"`
			RepostWithComment             bool        `json:"repostWithComment"`
			EnablePrivateChat             bool        `json:"enablePrivateChat"`
			BlockStrangerPoke             bool        `json:"blockStrangerPoke"`
			EnablePictureWatermark        bool        `json:"enablePictureWatermark"`
			EnableOperationPush           bool        `json:"enableOperationPush"`
			EnableChatSound               bool        `json:"enableChatSound"`
			ReplyNotificationAllowGroup   string      `json:"replyNotificationAllowGroup"`
			LikeNotificationOn            bool        `json:"likeNotificationOn"`
			RespectNotificationOn         bool        `json:"respectNotificationOn"`
			MentionNotificationOn         bool        `json:"mentionNotificationOn"`
			LiveNotificationOn            bool        `json:"liveNotificationOn"`
			OpenMessageTabOnLaunch        bool        `json:"openMessageTabOnLaunch"`
			TabOnLaunch                   string      `json:"tabOnLaunch"`
			HideSubscribeTab              interface{} `json:"hideSubscribeTab"`
			UndiscoverableByWeiboUser     interface{} `json:"undiscoverableByWeiboUser"`
			PaidMarket                    interface{} `json:"paidMarket"`
			FollowingUpdatesSortBy        string      `json:"followingUpdatesSortBy"`
			PersonalizedRecommendationOn  bool        `json:"personalizedRecommendationOn"`
		} `json:"preferences"`
		IsBetaUser         bool      `json:"isBetaUser"`
		UsernameSet        bool      `json:"usernameSet"`
		WechatOpenId       string    `json:"wechatOpenId"`
		WechatUnionId      string    `json:"wechatUnionId"`
		MobilePhoneNumber  string    `json:"mobilePhoneNumber"`
		AreaCode           string    `json:"areaCode"`
		GroupVersion       int       `json:"groupVersion"`
		LastChangeNameTime time.Time `json:"lastChangeNameTime"`
		Birthday           string    `json:"birthday"`
		WechatUserInfo     struct {
			Id           string `json:"id"`
			ExternalName string `json:"externalName"`
		} `json:"wechatUserInfo"`
		TopicRoles  []interface{} `json:"topicRoles"`
		Zodiac      string        `json:"zodiac"`
		ShowRespect bool          `json:"showRespect"`
		Medals      []interface{} `json:"medals"`
		ProfileTags []struct {
			PicUrl string `json:"picUrl,omitempty"`
			Type   string `json:"type"`
			Text   string `json:"text,omitempty"`
		} `json:"profileTags"`
		BackgroundMaskColor    string `json:"backgroundMaskColor"`
		IsLoginUser            bool   `json:"isLoginUser"`
		IsBanned               bool   `json:"isBanned"`
		UserHasPosted          bool   `json:"userHasPosted"`
		RegisterAppVersion     string `json:"registerAppVersion"`
		HasStories             bool   `json:"hasStories"`
		RestrictedAvatarChange struct {
			NextChangeDate string `json:"nextChangeDate"`
			Limits         int    `json:"limits"`
		} `json:"restrictedAvatarChange"`
		RestrictedNameChange struct {
			NextChangeDate string `json:"nextChangeDate"`
			Limits         int    `json:"limits"`
		} `json:"restrictedNameChange"`
		IsDefaultScreenName bool `json:"isDefaultScreenName"`
		ProfileVisitInfo    struct {
			TodayCount    int `json:"todayCount"`
			LatestVisitor struct {
				Id            string    `json:"id"`
				Username      string    `json:"username"`
				ScreenName    string    `json:"screenName"`
				CreatedAt     time.Time `json:"createdAt"`
				UpdatedAt     time.Time `json:"updatedAt"`
				IsVerified    bool      `json:"isVerified"`
				VerifyMessage string    `json:"verifyMessage"`
				BriefIntro    string    `json:"briefIntro"`
				AvatarImage   struct {
					ThumbnailUrl string `json:"thumbnailUrl"`
					SmallPicUrl  string `json:"smallPicUrl"`
					MiddlePicUrl string `json:"middlePicUrl"`
					PicUrl       string `json:"picUrl"`
					Format       string `json:"format"`
				} `json:"avatarImage"`
				Decorations struct {
				} `json:"decorations"`
				ProfileImageUrl string `json:"profileImageUrl"`
				StatsCount      struct {
					TopicSubscribed            int `json:"topicSubscribed"`
					TopicCreated               int `json:"topicCreated"`
					FollowedCount              int `json:"followedCount"`
					FollowingCount             int `json:"followingCount"`
					HighlightedPersonalUpdates int `json:"highlightedPersonalUpdates"`
					Liked                      int `json:"liked"`
					RespectedCount             int `json:"respectedCount"`
				} `json:"statsCount"`
				IsBannedForever bool `json:"isBannedForever"`
				IsSponsor       bool `json:"isSponsor"`
				BackgroundImage struct {
					PicUrl string `json:"picUrl"`
				} `json:"backgroundImage"`
			} `json:"latestVisitor"`
		} `json:"profileVisitInfo"`
	} `json:"user"`
	EnabledFeatures []string `json:"enabledFeatures"`
	AgreedProtocol  []string `json:"agreedProtocol"`
}

type HotFeedOutput struct {
	Data []struct {
		Id         string `json:"id"`
		Type       string `json:"type"`
		Content    string `json:"content"`
		UrlsInText []struct {
			Title       string `json:"title"`
			OriginalUrl string `json:"originalUrl"`
			Url         string `json:"url"`
			Type        string `json:"type"`
		} `json:"urlsInText"`
		Status             string `json:"status"`
		IsCommentForbidden bool   `json:"isCommentForbidden"`
		LikeCount          int    `json:"likeCount"`
		CommentCount       int    `json:"commentCount"`
		RepostCount        int    `json:"repostCount"`
		ShareCount         int    `json:"shareCount"`
		Topic              struct {
			Id               string `json:"id"`
			Type             string `json:"type"`
			Content          string `json:"content"`
			SubscribersCount int    `json:"subscribersCount"`
			SquarePicture    struct {
				Format       string      `json:"format"`
				PicUrl       string      `json:"picUrl"`
				MiddlePicUrl string      `json:"middlePicUrl"`
				SmallPicUrl  string      `json:"smallPicUrl"`
				ThumbnailUrl string      `json:"thumbnailUrl"`
				LivePhoto    interface{} `json:"livePhoto"`
				Themes       struct {
				} `json:"themes"`
				Nft interface{} `json:"nft"`
			} `json:"squarePicture"`
			BriefIntro               string        `json:"briefIntro"`
			TopicType                string        `json:"topicType"`
			OperateStatus            string        `json:"operateStatus"`
			IsValid                  bool          `json:"isValid"`
			IsVerified               bool          `json:"isVerified"`
			TopicId                  int           `json:"topicId"`
			IsSearchable             bool          `json:"isSearchable"`
			LikeIcon                 string        `json:"likeIcon"`
			MessagePrefix            string        `json:"messagePrefix"`
			InternalTags             []interface{} `json:"internalTags"`
			CustomTags               []interface{} `json:"customTags"`
			AuditStatus              string        `json:"auditStatus"`
			NewCategory              []interface{} `json:"newCategory"`
			InvolvedUsers            interface{}   `json:"involvedUsers"`
			EntryTab                 string        `json:"entryTab"`
			Tabs                     []interface{} `json:"tabs"`
			RectanglePicture         interface{}   `json:"rectanglePicture"`
			PictureUrl               string        `json:"pictureUrl"`
			ThumbnailUrl             string        `json:"thumbnailUrl"`
			SubscribedStatusRawValue int           `json:"subscribedStatusRawValue"`
			SubscribedAt             string        `json:"subscribedAt"`
			Ref                      string        `json:"ref"`
			IsDreamTopic             bool          `json:"isDreamTopic"`
			IsAnonymous              bool          `json:"isAnonymous"`
			EnablePictureComments    bool          `json:"enablePictureComments"`
			EnablePictureWatermark   bool          `json:"enablePictureWatermark"`
			TimeForRank              string        `json:"timeForRank"`
			LastMessagePostTime      time.Time     `json:"lastMessagePostTime"`
			CreatedAt                string        `json:"createdAt"`
			UpdatedAt                string        `json:"updatedAt"`
			SubscribersTextSuffix    string        `json:"subscribersTextSuffix"`
			SubscribersName          string        `json:"subscribersName"`
			FriendsAlsoSubscribe     string        `json:"friendsAlsoSubscribe"`
			Maintainer               interface{}   `json:"maintainer"`
			IsUserTopicAdmin         bool          `json:"isUserTopicAdmin"`
			ActivitySection          interface{}   `json:"activitySection"`
			ActivitySections         []interface{} `json:"activitySections"`
			Tips                     struct {
				InDraft   string `json:"inDraft"`
				InComment string `json:"inComment"`
			} `json:"tips"`
			ToppingArea                 interface{}   `json:"toppingArea"`
			InShortcuts                 bool          `json:"inShortcuts"`
			PreferSection               string        `json:"preferSection"`
			RelatedHashtags             []interface{} `json:"relatedHashtags"`
			Intro                       string        `json:"intro"`
			SquarePostUpdateTime        time.Time     `json:"squarePostUpdateTime"`
			SubscribersAction           string        `json:"subscribersAction"`
			ApproximateSubscribersCount string        `json:"approximateSubscribersCount"`
			SubscribersDescription      string        `json:"subscribersDescription"`
			IsCommentForbidden          bool          `json:"isCommentForbidden"`
			BotCount                    int           `json:"botCount"`
			RecentPost                  string        `json:"recentPost"`
			Source                      string        `json:"source"`
			EnableForUserPost           bool          `json:"enableForUserPost"`
		} `json:"topic"`
		Pictures []struct {
			Key             string  `json:"key"`
			ThumbnailUrl    string  `json:"thumbnailUrl"`
			SmallPicUrl     string  `json:"smallPicUrl"`
			MiddlePicUrl    string  `json:"middlePicUrl"`
			PicUrl          string  `json:"picUrl"`
			Format          string  `json:"format"`
			CropperPosX     float64 `json:"cropperPosX"`
			CropperPosY     float64 `json:"cropperPosY"`
			Width           int     `json:"width"`
			Height          int     `json:"height"`
			WatermarkPicUrl string  `json:"watermarkPicUrl"`
		} `json:"pictures"`
		Collected   bool        `json:"collected"`
		CollectTime interface{} `json:"collectTime"`
		User        struct {
			Id            string    `json:"id"`
			Username      string    `json:"username"`
			ScreenName    string    `json:"screenName"`
			CreatedAt     time.Time `json:"createdAt"`
			UpdatedAt     time.Time `json:"updatedAt"`
			IsVerified    bool      `json:"isVerified"`
			VerifyMessage string    `json:"verifyMessage"`
			BriefIntro    string    `json:"briefIntro"`
			AvatarImage   struct {
				ThumbnailUrl string `json:"thumbnailUrl"`
				SmallPicUrl  string `json:"smallPicUrl"`
				MiddlePicUrl string `json:"middlePicUrl"`
				PicUrl       string `json:"picUrl"`
				Format       string `json:"format"`
			} `json:"avatarImage"`
			Decorations struct {
				Sponsor struct {
					PicUrl string `json:"picUrl"`
					Themes struct {
						Dark struct {
							PicUrl string `json:"picUrl"`
						} `json:"dark"`
					} `json:"themes"`
				} `json:"sponsor,omitempty"`
			} `json:"decorations"`
			ProfileImageUrl string `json:"profileImageUrl"`
			StatsCount      struct {
				TopicSubscribed            int `json:"topicSubscribed"`
				TopicCreated               int `json:"topicCreated"`
				FollowedCount              int `json:"followedCount"`
				FollowingCount             int `json:"followingCount"`
				HighlightedPersonalUpdates int `json:"highlightedPersonalUpdates"`
				Liked                      int `json:"liked"`
				RespectedCount             int `json:"respectedCount"`
			} `json:"statsCount"`
			IsBannedForever  bool      `json:"isBannedForever"`
			IsSponsor        bool      `json:"isSponsor"`
			SponsorExpiresAt time.Time `json:"sponsorExpiresAt,omitempty"`
			BackgroundImage  struct {
				PicUrl string `json:"picUrl"`
			} `json:"backgroundImage"`
			Bio       string `json:"bio"`
			Gender    string `json:"gender"`
			City      string `json:"city,omitempty"`
			Country   string `json:"country,omitempty"`
			Province  string `json:"province,omitempty"`
			RefRemark struct {
				Type  string `json:"type"`
				RefId string `json:"refId"`
			} `json:"refRemark"`
			Ref           string `json:"ref"`
			StoryStatus   string `json:"storyStatus"`
			TrailingIcons []struct {
				Picture struct {
					PicUrl string `json:"picUrl"`
					Format string `json:"format"`
				} `json:"picture"`
				PicUrl string `json:"picUrl"`
				Url    string `json:"url"`
			} `json:"trailingIcons,omitempty"`
		} `json:"user"`
		CreatedAt             time.Time `json:"createdAt"`
		IsFeatured            bool      `json:"isFeatured"`
		EnablePictureComments bool      `json:"enablePictureComments"`
		Repostable            bool      `json:"repostable"`
		Rollouts              struct {
			WmpShare struct {
				Id   string `json:"id"`
				Path string `json:"path"`
			} `json:"wmpShare"`
		} `json:"rollouts"`
		Subtitle           string `json:"subtitle"`
		ScrollingSubtitles []struct {
			Text string `json:"text"`
			Type string `json:"type"`
		} `json:"scrollingSubtitles"`
		ReadTrackInfo struct {
			StoryStatus string `json:"storyStatus"`
		} `json:"readTrackInfo"`
		LikeIcon string `json:"likeIcon,omitempty"`
	} `json:"data"`
	LoadMoreKey HotFeedLoadMoreKey `json:"loadMoreKey"`
}

type HotFeedLoadMoreKey struct {
	Offset int `json:"offset"`
}

type TimelineOutput struct {
	Success bool `json:"success"`
	Data    []struct {
		ActionTime time.Time `json:"actionTime"`
		Id         string    `json:"id"`
		Type       string    `json:"type"`
		Content    string    `json:"content,omitempty"`
		UrlsInText []struct {
			Title       string `json:"title"`
			OriginalUrl string `json:"originalUrl"`
			Url         string `json:"url"`
			Type        string `json:"type"`
		} `json:"urlsInText,omitempty"`
		Status             string `json:"status,omitempty"`
		IsCommentForbidden bool   `json:"isCommentForbidden,omitempty"`
		LikeCount          int    `json:"likeCount,omitempty"`
		CommentCount       int    `json:"commentCount,omitempty"`
		RepostCount        int    `json:"repostCount,omitempty"`
		ShareCount         int    `json:"shareCount,omitempty"`
		Topic              struct {
			Id               string `json:"id"`
			Type             string `json:"type"`
			Content          string `json:"content"`
			SubscribersCount int    `json:"subscribersCount"`
			SquarePicture    struct {
				Format       string      `json:"format"`
				PicUrl       string      `json:"picUrl"`
				MiddlePicUrl string      `json:"middlePicUrl"`
				SmallPicUrl  string      `json:"smallPicUrl"`
				ThumbnailUrl string      `json:"thumbnailUrl"`
				LivePhoto    interface{} `json:"livePhoto"`
				Themes       struct {
				} `json:"themes"`
				Nft interface{} `json:"nft"`
			} `json:"squarePicture"`
			BriefIntro               string        `json:"briefIntro"`
			TopicType                string        `json:"topicType"`
			OperateStatus            string        `json:"operateStatus"`
			IsValid                  bool          `json:"isValid"`
			IsVerified               bool          `json:"isVerified"`
			TopicId                  int           `json:"topicId"`
			IsSearchable             bool          `json:"isSearchable"`
			LikeIcon                 string        `json:"likeIcon"`
			MessagePrefix            string        `json:"messagePrefix"`
			InternalTags             []interface{} `json:"internalTags"`
			CustomTags               []interface{} `json:"customTags"`
			AuditStatus              string        `json:"auditStatus"`
			NewCategory              []interface{} `json:"newCategory"`
			InvolvedUsers            interface{}   `json:"involvedUsers"`
			EntryTab                 string        `json:"entryTab"`
			Tabs                     []interface{} `json:"tabs"`
			RectanglePicture         interface{}   `json:"rectanglePicture"`
			PictureUrl               string        `json:"pictureUrl"`
			ThumbnailUrl             string        `json:"thumbnailUrl"`
			SubscribedStatusRawValue int           `json:"subscribedStatusRawValue"`
			SubscribedAt             string        `json:"subscribedAt"`
			Ref                      string        `json:"ref"`
			IsDreamTopic             bool          `json:"isDreamTopic"`
			IsAnonymous              bool          `json:"isAnonymous"`
			EnablePictureComments    bool          `json:"enablePictureComments"`
			EnablePictureWatermark   bool          `json:"enablePictureWatermark"`
			TimeForRank              string        `json:"timeForRank"`
			LastMessagePostTime      time.Time     `json:"lastMessagePostTime"`
			CreatedAt                string        `json:"createdAt"`
			UpdatedAt                string        `json:"updatedAt"`
			SubscribersTextSuffix    string        `json:"subscribersTextSuffix"`
			SubscribersName          string        `json:"subscribersName"`
			FriendsAlsoSubscribe     string        `json:"friendsAlsoSubscribe"`
			Maintainer               interface{}   `json:"maintainer"`
			IsUserTopicAdmin         bool          `json:"isUserTopicAdmin"`
			ActivitySection          interface{}   `json:"activitySection"`
			ActivitySections         []interface{} `json:"activitySections"`
			Tips                     struct {
				InDraft   string `json:"inDraft"`
				InComment string `json:"inComment"`
			} `json:"tips"`
			ToppingArea                 interface{}   `json:"toppingArea"`
			InShortcuts                 bool          `json:"inShortcuts"`
			PreferSection               string        `json:"preferSection"`
			RelatedHashtags             []interface{} `json:"relatedHashtags"`
			Intro                       string        `json:"intro"`
			SquarePostUpdateTime        time.Time     `json:"squarePostUpdateTime"`
			SubscribersAction           string        `json:"subscribersAction"`
			ApproximateSubscribersCount string        `json:"approximateSubscribersCount"`
			SubscribersDescription      string        `json:"subscribersDescription"`
			IsCommentForbidden          bool          `json:"isCommentForbidden"`
			BotCount                    int           `json:"botCount"`
			RecentPost                  string        `json:"recentPost"`
			Source                      string        `json:"source"`
			EnableForUserPost           bool          `json:"enableForUserPost"`
		} `json:"topic,omitempty"`
		Pictures []struct {
			Key             string  `json:"key"`
			ThumbnailUrl    string  `json:"thumbnailUrl"`
			SmallPicUrl     string  `json:"smallPicUrl"`
			MiddlePicUrl    string  `json:"middlePicUrl"`
			PicUrl          string  `json:"picUrl"`
			Format          string  `json:"format"`
			CropperPosX     float64 `json:"cropperPosX"`
			CropperPosY     float64 `json:"cropperPosY"`
			Width           int     `json:"width"`
			Height          int     `json:"height"`
			WatermarkPicUrl string  `json:"watermarkPicUrl,omitempty"`
		} `json:"pictures,omitempty"`
		Collected   bool        `json:"collected,omitempty"`
		CollectTime interface{} `json:"collectTime"`
		User        struct {
			Id            string    `json:"id"`
			Username      string    `json:"username"`
			ScreenName    string    `json:"screenName"`
			CreatedAt     time.Time `json:"createdAt"`
			UpdatedAt     time.Time `json:"updatedAt"`
			IsVerified    bool      `json:"isVerified"`
			VerifyMessage string    `json:"verifyMessage"`
			BriefIntro    string    `json:"briefIntro"`
			AvatarImage   struct {
				ThumbnailUrl string `json:"thumbnailUrl"`
				SmallPicUrl  string `json:"smallPicUrl"`
				MiddlePicUrl string `json:"middlePicUrl"`
				PicUrl       string `json:"picUrl"`
				Format       string `json:"format"`
				Nft          struct {
					Name       string `json:"name"`
					Collection string `json:"collection"`
					OpenseaURL string `json:"openseaURL"`
					AssetURL   string `json:"assetURL"`
				} `json:"nft,omitempty"`
			} `json:"avatarImage"`
			Decorations struct {
				Sponsor struct {
					PicUrl string `json:"picUrl"`
					Themes struct {
						Dark struct {
							PicUrl string `json:"picUrl"`
						} `json:"dark"`
					} `json:"themes"`
				} `json:"sponsor,omitempty"`
			} `json:"decorations"`
			ProfileImageUrl string `json:"profileImageUrl"`
			StatsCount      struct {
				TopicSubscribed            int `json:"topicSubscribed"`
				TopicCreated               int `json:"topicCreated"`
				FollowedCount              int `json:"followedCount"`
				FollowingCount             int `json:"followingCount"`
				HighlightedPersonalUpdates int `json:"highlightedPersonalUpdates"`
				Liked                      int `json:"liked"`
				RespectedCount             int `json:"respectedCount"`
			} `json:"statsCount"`
			IsBannedForever  bool      `json:"isBannedForever"`
			IsSponsor        bool      `json:"isSponsor"`
			SponsorExpiresAt time.Time `json:"sponsorExpiresAt,omitempty"`
			TrailingIcons    []struct {
				Picture struct {
					PicUrl string `json:"picUrl"`
					Format string `json:"format"`
				} `json:"picture"`
				PicUrl string `json:"picUrl"`
				Url    string `json:"url"`
			} `json:"trailingIcons,omitempty"`
			BackgroundImage struct {
				PicUrl string `json:"picUrl"`
			} `json:"backgroundImage,omitempty"`
			Bio       string `json:"bio"`
			Gender    string `json:"gender"`
			City      string `json:"city"`
			Country   string `json:"country"`
			Province  string `json:"province"`
			RefRemark struct {
				Type  string `json:"type"`
				RefId string `json:"refId"`
			} `json:"refRemark"`
			StoryStatus string `json:"storyStatus"`
			Following   bool   `json:"following,omitempty"`
		} `json:"user,omitempty"`
		CreatedAt             time.Time `json:"createdAt"`
		IsFeatured            bool      `json:"isFeatured,omitempty"`
		EnablePictureComments bool      `json:"enablePictureComments,omitempty"`
		Repostable            bool      `json:"repostable,omitempty"`
		Rollouts              struct {
			WmpShare struct {
				Id   string `json:"id"`
				Path string `json:"path"`
			} `json:"wmpShare"`
		} `json:"rollouts,omitempty"`
		ScrollingSubtitles []struct {
			Text string `json:"text"`
			Type string `json:"type"`
		} `json:"scrollingSubtitles,omitempty"`
		LikeIcon      string `json:"likeIcon,omitempty"`
		ReadTrackInfo struct {
			StoryStatus string `json:"storyStatus,omitempty"`
			LoadedAt    int64  `json:"loadedAt"`
			FeedType    string `json:"feedType"`
		} `json:"readTrackInfo"`
		RawContent string `json:"rawContent,omitempty"`
		Target     struct {
			Type       string        `json:"type"`
			Id         string        `json:"id"`
			Content    string        `json:"content"`
			RawContent string        `json:"rawContent,omitempty"`
			UrlsInText []interface{} `json:"urlsInText"`
			Status     string        `json:"status"`
			User       struct {
				Id            string    `json:"id"`
				Username      string    `json:"username"`
				ScreenName    string    `json:"screenName"`
				CreatedAt     time.Time `json:"createdAt"`
				UpdatedAt     time.Time `json:"updatedAt"`
				IsVerified    bool      `json:"isVerified"`
				VerifyMessage string    `json:"verifyMessage"`
				BriefIntro    string    `json:"briefIntro"`
				AvatarImage   struct {
					ThumbnailUrl string `json:"thumbnailUrl"`
					SmallPicUrl  string `json:"smallPicUrl"`
					MiddlePicUrl string `json:"middlePicUrl"`
					PicUrl       string `json:"picUrl"`
					Format       string `json:"format"`
					Nft          struct {
						Name       string `json:"name"`
						Collection string `json:"collection"`
						OpenseaURL string `json:"openseaURL"`
						AssetURL   string `json:"assetURL"`
					} `json:"nft,omitempty"`
				} `json:"avatarImage"`
				Decorations struct {
					Sponsor struct {
						PicUrl string `json:"picUrl"`
						Themes struct {
							Dark struct {
								PicUrl string `json:"picUrl"`
							} `json:"dark"`
						} `json:"themes"`
					} `json:"sponsor,omitempty"`
				} `json:"decorations"`
				ProfileImageUrl string `json:"profileImageUrl"`
				StatsCount      struct {
					TopicSubscribed            int `json:"topicSubscribed"`
					TopicCreated               int `json:"topicCreated"`
					FollowedCount              int `json:"followedCount"`
					FollowingCount             int `json:"followingCount"`
					HighlightedPersonalUpdates int `json:"highlightedPersonalUpdates"`
					Liked                      int `json:"liked"`
					RespectedCount             int `json:"respectedCount"`
				} `json:"statsCount"`
				IsBannedForever bool `json:"isBannedForever"`
				IsSponsor       bool `json:"isSponsor"`
				BackgroundImage struct {
					PicUrl string `json:"picUrl"`
				} `json:"backgroundImage"`
				Bio              string    `json:"bio"`
				City             string    `json:"city"`
				Country          string    `json:"country"`
				Province         string    `json:"province"`
				Ref              string    `json:"ref"`
				SponsorExpiresAt time.Time `json:"sponsorExpiresAt,omitempty"`
				Gender           string    `json:"gender,omitempty"`
				RefRemark        struct {
					Type  string `json:"type"`
					RefId string `json:"refId"`
				} `json:"refRemark,omitempty"`
				TrailingIcons []struct {
					Picture struct {
						PicUrl string `json:"picUrl"`
						Format string `json:"format"`
					} `json:"picture"`
					PicUrl string `json:"picUrl"`
					Url    string `json:"url"`
				} `json:"trailingIcons,omitempty"`
			} `json:"user"`
			CommentCount int    `json:"commentCount"`
			RepostCount  int    `json:"repostCount"`
			LikeCount    int    `json:"likeCount"`
			ShareCount   int    `json:"shareCount"`
			RootType     string `json:"rootType,omitempty"`
			Pictures     []struct {
				Key             string  `json:"key"`
				ThumbnailUrl    string  `json:"thumbnailUrl"`
				SmallPicUrl     string  `json:"smallPicUrl"`
				MiddlePicUrl    string  `json:"middlePicUrl"`
				PicUrl          string  `json:"picUrl"`
				Format          string  `json:"format"`
				CropperPosX     float64 `json:"cropperPosX"`
				CropperPosY     float64 `json:"cropperPosY"`
				Width           int     `json:"width"`
				Height          int     `json:"height"`
				WatermarkPicUrl string  `json:"watermarkPicUrl"`
			} `json:"pictures"`
			CreatedAt          time.Time `json:"createdAt"`
			Liked              bool      `json:"liked,omitempty"`
			Collected          bool      `json:"collected"`
			IsCommentForbidden bool      `json:"isCommentForbidden,omitempty"`
			Topic              struct {
				Id               string `json:"id"`
				Type             string `json:"type"`
				Content          string `json:"content"`
				SubscribersCount int    `json:"subscribersCount"`
				SquarePicture    struct {
					Format       string      `json:"format"`
					PicUrl       string      `json:"picUrl"`
					MiddlePicUrl string      `json:"middlePicUrl"`
					SmallPicUrl  string      `json:"smallPicUrl"`
					ThumbnailUrl string      `json:"thumbnailUrl"`
					LivePhoto    interface{} `json:"livePhoto"`
					Themes       struct {
					} `json:"themes"`
					Nft interface{} `json:"nft"`
				} `json:"squarePicture"`
				BriefIntro               string        `json:"briefIntro"`
				TopicType                string        `json:"topicType"`
				OperateStatus            string        `json:"operateStatus"`
				IsValid                  bool          `json:"isValid"`
				IsVerified               bool          `json:"isVerified"`
				TopicId                  int           `json:"topicId"`
				IsSearchable             bool          `json:"isSearchable"`
				LikeIcon                 string        `json:"likeIcon"`
				MessagePrefix            string        `json:"messagePrefix"`
				InternalTags             []interface{} `json:"internalTags"`
				CustomTags               []interface{} `json:"customTags"`
				AuditStatus              string        `json:"auditStatus"`
				NewCategory              []interface{} `json:"newCategory"`
				InvolvedUsers            interface{}   `json:"involvedUsers"`
				EntryTab                 string        `json:"entryTab"`
				Tabs                     []interface{} `json:"tabs"`
				RectanglePicture         interface{}   `json:"rectanglePicture"`
				PictureUrl               string        `json:"pictureUrl"`
				ThumbnailUrl             string        `json:"thumbnailUrl"`
				SubscribedStatusRawValue int           `json:"subscribedStatusRawValue"`
				SubscribedAt             string        `json:"subscribedAt"`
				Ref                      string        `json:"ref"`
				IsDreamTopic             bool          `json:"isDreamTopic"`
				IsAnonymous              bool          `json:"isAnonymous"`
				EnablePictureComments    bool          `json:"enablePictureComments"`
				EnablePictureWatermark   bool          `json:"enablePictureWatermark"`
				TimeForRank              string        `json:"timeForRank"`
				LastMessagePostTime      time.Time     `json:"lastMessagePostTime"`
				CreatedAt                string        `json:"createdAt"`
				UpdatedAt                string        `json:"updatedAt"`
				SubscribersTextSuffix    string        `json:"subscribersTextSuffix"`
				SubscribersName          string        `json:"subscribersName"`
				FriendsAlsoSubscribe     string        `json:"friendsAlsoSubscribe"`
				Maintainer               interface{}   `json:"maintainer"`
				IsUserTopicAdmin         bool          `json:"isUserTopicAdmin"`
				ActivitySection          interface{}   `json:"activitySection"`
				ActivitySections         []interface{} `json:"activitySections"`
				Tips                     struct {
					InDraft   string `json:"inDraft"`
					InComment string `json:"inComment"`
				} `json:"tips"`
				ToppingArea                 interface{}   `json:"toppingArea"`
				InShortcuts                 bool          `json:"inShortcuts"`
				PreferSection               string        `json:"preferSection"`
				RelatedHashtags             []interface{} `json:"relatedHashtags"`
				Intro                       string        `json:"intro"`
				SquarePostUpdateTime        time.Time     `json:"squarePostUpdateTime"`
				SubscribersAction           string        `json:"subscribersAction"`
				ApproximateSubscribersCount string        `json:"approximateSubscribersCount"`
				SubscribersDescription      string        `json:"subscribersDescription"`
				IsCommentForbidden          bool          `json:"isCommentForbidden"`
				BotCount                    int           `json:"botCount"`
				RecentPost                  string        `json:"recentPost"`
				Source                      string        `json:"source"`
				EnableForUserPost           bool          `json:"enableForUserPost"`
			} `json:"topic,omitempty"`
			CollectTime           interface{} `json:"collectTime"`
			IsFeatured            bool        `json:"isFeatured,omitempty"`
			EnablePictureComments bool        `json:"enablePictureComments,omitempty"`
			Repostable            bool        `json:"repostable,omitempty"`
			Rollouts              struct {
				WmpShare struct {
					Id   string `json:"id"`
					Path string `json:"path"`
				} `json:"wmpShare"`
			} `json:"rollouts,omitempty"`
			Subtitle           string `json:"subtitle,omitempty"`
			ScrollingSubtitles []struct {
				Text string `json:"text"`
				Type string `json:"type"`
			} `json:"scrollingSubtitles,omitempty"`
			LikeIcon string    `json:"likeIcon,omitempty"`
			EditedAt time.Time `json:"editedAt,omitempty"`
			RecordId string    `json:"recordId,omitempty"`
		} `json:"target,omitempty"`
		TargetType      string   `json:"targetType,omitempty"`
		RootType        string   `json:"rootType,omitempty"`
		Liked           bool     `json:"liked,omitempty"`
		UpdateIds       []string `json:"updateIds,omitempty"`
		Action          string   `json:"action,omitempty"`
		Usernames       []string `json:"usernames,omitempty"`
		TargetUsernames []string `json:"targetUsernames,omitempty"`
		Users           []struct {
			Id            string    `json:"id"`
			Username      string    `json:"username"`
			ScreenName    string    `json:"screenName"`
			CreatedAt     time.Time `json:"createdAt"`
			UpdatedAt     time.Time `json:"updatedAt"`
			IsVerified    bool      `json:"isVerified"`
			VerifyMessage string    `json:"verifyMessage"`
			BriefIntro    string    `json:"briefIntro"`
			AvatarImage   struct {
				ThumbnailUrl string `json:"thumbnailUrl"`
				SmallPicUrl  string `json:"smallPicUrl"`
				MiddlePicUrl string `json:"middlePicUrl"`
				PicUrl       string `json:"picUrl"`
				Format       string `json:"format"`
				Nft          struct {
					Name       string `json:"name"`
					Collection string `json:"collection"`
					OpenseaURL string `json:"openseaURL"`
					AssetURL   string `json:"assetURL"`
				} `json:"nft,omitempty"`
			} `json:"avatarImage"`
			Decorations struct {
				Sponsor struct {
					PicUrl string `json:"picUrl"`
					Themes struct {
						Dark struct {
							PicUrl string `json:"picUrl"`
						} `json:"dark"`
					} `json:"themes"`
				} `json:"sponsor"`
			} `json:"decorations"`
			ProfileImageUrl string `json:"profileImageUrl"`
			StatsCount      struct {
				TopicSubscribed            int `json:"topicSubscribed"`
				TopicCreated               int `json:"topicCreated"`
				FollowedCount              int `json:"followedCount"`
				FollowingCount             int `json:"followingCount"`
				HighlightedPersonalUpdates int `json:"highlightedPersonalUpdates"`
				Liked                      int `json:"liked"`
				RespectedCount             int `json:"respectedCount"`
			} `json:"statsCount"`
			IsBannedForever  bool      `json:"isBannedForever"`
			IsSponsor        bool      `json:"isSponsor"`
			SponsorExpiresAt time.Time `json:"sponsorExpiresAt"`
			TrailingIcons    []struct {
				Picture struct {
					PicUrl string `json:"picUrl"`
					Format string `json:"format"`
				} `json:"picture"`
				PicUrl string `json:"picUrl"`
				Url    string `json:"url"`
			} `json:"trailingIcons,omitempty"`
			BackgroundImage struct {
				PicUrl string `json:"picUrl"`
			} `json:"backgroundImage,omitempty"`
			Bio         string `json:"bio"`
			Gender      string `json:"gender"`
			City        string `json:"city"`
			Country     string `json:"country"`
			Province    string `json:"province"`
			Ref         string `json:"ref,omitempty"`
			StoryStatus string `json:"storyStatus"`
			Following   bool   `json:"following,omitempty"`
		} `json:"users,omitempty"`
		TargetUsers []struct {
			Id            string    `json:"id"`
			Username      string    `json:"username"`
			ScreenName    string    `json:"screenName"`
			CreatedAt     time.Time `json:"createdAt"`
			UpdatedAt     time.Time `json:"updatedAt"`
			IsVerified    bool      `json:"isVerified"`
			VerifyMessage string    `json:"verifyMessage"`
			BriefIntro    string    `json:"briefIntro"`
			AvatarImage   struct {
				ThumbnailUrl string `json:"thumbnailUrl"`
				SmallPicUrl  string `json:"smallPicUrl"`
				MiddlePicUrl string `json:"middlePicUrl"`
				PicUrl       string `json:"picUrl"`
				Format       string `json:"format"`
			} `json:"avatarImage"`
			Decorations struct {
				Sponsor struct {
					PicUrl string `json:"picUrl"`
					Themes struct {
						Dark struct {
							PicUrl string `json:"picUrl"`
						} `json:"dark"`
					} `json:"themes"`
				} `json:"sponsor,omitempty"`
			} `json:"decorations"`
			ProfileImageUrl string `json:"profileImageUrl"`
			StatsCount      struct {
				TopicSubscribed            int `json:"topicSubscribed"`
				TopicCreated               int `json:"topicCreated"`
				FollowedCount              int `json:"followedCount"`
				FollowingCount             int `json:"followingCount"`
				HighlightedPersonalUpdates int `json:"highlightedPersonalUpdates"`
				Liked                      int `json:"liked"`
				RespectedCount             int `json:"respectedCount"`
			} `json:"statsCount"`
			IsBannedForever bool `json:"isBannedForever"`
			IsSponsor       bool `json:"isSponsor"`
			BackgroundImage struct {
				PicUrl string `json:"picUrl"`
			} `json:"backgroundImage,omitempty"`
			Bio              string    `json:"bio,omitempty"`
			City             string    `json:"city,omitempty"`
			Country          string    `json:"country,omitempty"`
			Province         string    `json:"province,omitempty"`
			Ref              string    `json:"ref"`
			Gender           string    `json:"gender,omitempty"`
			SponsorExpiresAt time.Time `json:"sponsorExpiresAt,omitempty"`
			Following        bool      `json:"following,omitempty"`
		} `json:"targetUsers,omitempty"`
		SyncCommentId  string `json:"syncCommentId,omitempty"`
		ReplyToComment struct {
			Type       string        `json:"type"`
			Id         string        `json:"id"`
			TargetType string        `json:"targetType"`
			TargetId   string        `json:"targetId"`
			ThreadId   string        `json:"threadId"`
			CreatedAt  time.Time     `json:"createdAt"`
			Level      int           `json:"level"`
			Content    string        `json:"content"`
			UrlsInText []interface{} `json:"urlsInText"`
			LikeCount  int           `json:"likeCount"`
			ReplyCount int           `json:"replyCount"`
			Status     string        `json:"status"`
			User       struct {
				Id            string    `json:"id"`
				Username      string    `json:"username"`
				ScreenName    string    `json:"screenName"`
				CreatedAt     time.Time `json:"createdAt"`
				UpdatedAt     time.Time `json:"updatedAt"`
				IsVerified    bool      `json:"isVerified"`
				VerifyMessage string    `json:"verifyMessage"`
				BriefIntro    string    `json:"briefIntro"`
				AvatarImage   struct {
					ThumbnailUrl string `json:"thumbnailUrl"`
					SmallPicUrl  string `json:"smallPicUrl"`
					MiddlePicUrl string `json:"middlePicUrl"`
					PicUrl       string `json:"picUrl"`
					Format       string `json:"format"`
				} `json:"avatarImage"`
				Decorations struct {
					Sponsor struct {
						PicUrl string `json:"picUrl"`
						Themes struct {
							Dark struct {
								PicUrl string `json:"picUrl"`
							} `json:"dark"`
						} `json:"themes"`
					} `json:"sponsor"`
				} `json:"decorations"`
				ProfileImageUrl string `json:"profileImageUrl"`
				StatsCount      struct {
					TopicSubscribed            int `json:"topicSubscribed"`
					TopicCreated               int `json:"topicCreated"`
					FollowedCount              int `json:"followedCount"`
					FollowingCount             int `json:"followingCount"`
					HighlightedPersonalUpdates int `json:"highlightedPersonalUpdates"`
					Liked                      int `json:"liked"`
					RespectedCount             int `json:"respectedCount"`
				} `json:"statsCount"`
				IsBannedForever  bool      `json:"isBannedForever"`
				IsSponsor        bool      `json:"isSponsor"`
				SponsorExpiresAt time.Time `json:"sponsorExpiresAt"`
				BackgroundImage  struct {
					PicUrl string `json:"picUrl"`
				} `json:"backgroundImage"`
				Bio      string `json:"bio"`
				Gender   string `json:"gender"`
				City     string `json:"city"`
				Country  string `json:"country"`
				Province string `json:"province"`
			} `json:"user"`
			Pictures              []interface{} `json:"pictures"`
			Liked                 bool          `json:"liked"`
			EnablePictureComments bool          `json:"enablePictureComments"`
			ReadTrackInfo         struct {
			} `json:"readTrackInfo"`
			Collapsed   bool `json:"collapsed"`
			Collapsible bool `json:"collapsible"`
			IsPinned    bool `json:"isPinned"`
			Pinnable    bool `json:"pinnable"`
		} `json:"replyToComment,omitempty"`
	} `json:"data"`
	LoadMoreKey TimelineLoadMoreKey `json:"loadMoreKey"`
}

type TimelineLoadMoreKey struct {
	Session              string `json:"session"`
	LastPageEarliestTime int64  `json:"lastPageEarliestTime"`
	LastReadTime         int64  `json:"lastReadTime"`
}
