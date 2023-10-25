package tootecho

import (
	"github.com/benpate/toot"
	"github.com/benpate/toot/route"
	"github.com/benpate/toot/scope"
	"github.com/labstack/echo/v4"
)

type echoMethod func(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route

func Register[AuthToken toot.ScopesGetter](e *echo.Echo, api toot.API[AuthToken], middleware ...echo.MiddlewareFunc) {

	// https://docs.joinmastodon.org/methods/accounts/
	register(api, e.POST, route.PostAccount, api.PostAccount, scope.PostAccount)
	register(api, e.GET, route.GetAccount_VerifyCredentials, api.GetAccount_VerifyCredentials, scope.GetAccount_VerifyCredentials)
	register(api, e.PATCH, route.PatchAccount_UpdateCredentials, api.PatchAccount_UpdateCredentials, scope.PatchAccount_UpdateCredentials)
	register(api, e.GET, route.GetAccount, api.GetAccount, scope.GetAccount)
	register(api, e.GET, route.GetAccount_Statuses, api.GetAccount_Statuses, scope.GetAccount_Statuses)
	register(api, e.GET, route.GetAccount_Followers, api.GetAccount_Followers, scope.GetAccount_Followers)
	register(api, e.GET, route.GetAccount_Following, api.GetAccount_Following, scope.GetAccount_Following)
	register(api, e.GET, route.GetAccount_FeaturedTags, api.GetAccount_FeaturedTags, scope.GetAccount_FeaturedTags)
	register(api, e.POST, route.PostAccount, api.PostAccount_Follow, scope.PostAccont_Follow)
	register(api, e.POST, route.PostAccount_Unfollow, api.PostAccount_Unfollow, scope.PostAccount_Unfollow)
	register(api, e.POST, route.PostAccount_Block, api.PostAccount_Block, scope.PostAccount_Block)
	register(api, e.POST, route.PostAccount_Unblock, api.PostAccount_Unblock, scope.PostAccount_Unblock)
	register(api, e.POST, route.PostAccount_Mute, api.PostAccount_Mute, scope.PostAccount_Mute)
	register(api, e.POST, route.PostAccount_Unmute, api.PostAccount_Unmute, scope.PostAccount_Unmute)
	register(api, e.POST, route.PostAccount_Pin, api.PostAccount_Pin, scope.PostAccount_Pin)
	register(api, e.POST, route.PostAccount_Unpin, api.PostAccount_Unpin, scope.PostAccount_Unpin)
	register(api, e.POST, route.PostAccount_Note, api.PostAccount_Note, scope.PostAccount_Note)
	register(api, e.GET, route.PostAccount_Relationships, api.GetAccount_Relationships, scope.PostAccount_Relationships)
	register(api, e.GET, route.GetAccount_FamiliarFollowers, api.GetAccount_FamiliarFollowers, scope.GetAccount_FamiliarFollowers)
	register(api, e.GET, route.GetAccount_Search, api.GetAccount_Search, scope.GetAccount_Search)
	register(api, e.GET, route.GetAccount_Lookup, api.GetAccount_Lookup, scope.GetAccount_Lookup)

	// https://docs.joinmastodon.org/methods/apps/
	register(api, e.POST, route.PostApplication, api.PostApplication, scope.PostApplication)
	register(api, e.GET, route.GetApplication_VerifyCredentials, api.GetApplication_VerifyCredentials, scope.GetApplication_VerifyCredentials)

	// https://docs.joinmastodon.org/methods/announcements/
	register(api, e.GET, route.GetAnnouncements, api.GetAnnouncements, scope.GetAnnouncements)
	register(api, e.POST, route.PostAnnouncement_Dismiss, api.PostAnnouncement_Dismiss, scope.PostAnnoucement_Dismis)
	register(api, e.PUT, route.PutAnnouncement_Reaction, api.PutAnnouncement_Reaction, scope.PutAnnouncement_Reaction)
	register(api, e.DELETE, route.DeleteAnnouncement_Reaction, api.DeleteAnnouncement_Reaction, scope.DeleteAnnouncement_Reaction)

	// https://docs.joinmastodon.org/methods/apps/
	register(api, e.POST, route.PostApplication, api.PostApplication, scope.PostApplication)
	register(api, e.GET, route.GetApplication_VerifyCredentials, api.GetApplication_VerifyCredentials, scope.GetApplication_VerifyCredentials)

	// https://docs.joinmastodon.org/methods/blocks/
	register(api, e.GET, route.GetBlocks, api.GetBlocks, scope.GetBlocks)

	// https://docs.joinmastodon.org/methods/bookmarks/
	register(api, e.GET, route.GetBookmarks, api.GetBookmarks, scope.GetBookmarks)

	// https://docs.joinmastodon.org/methods/conversations/
	register(api, e.GET, route.GetConversations, api.GetConversations, scope.GetConversations)
	register(api, e.DELETE, route.DeleteConversation, api.DeleteConversation, scope.DeleteConversation)
	register(api, e.POST, route.PostConversationRead, api.PostConversationRead, scope.PostConversationRead)

	// https://docs.joinmastodon.org/methods/custom_emojis/
	register(api, e.GET, route.GetCustomEmojis, api.GetCustomEmojis, scope.GetCustomEmojis)

	// https://docs.joinmastodon.org/methods/directory/
	register(api, e.GET, route.GetDirectory, api.GetDirectory, scope.GetDirectory)

	// https://docs.joinmastodon.org/methods/domain_blocks/
	register(api, e.GET, route.GetDomainBlocks, api.GetDomainBlocks, scope.GetDomainBlocks)
	register(api, e.POST, route.PostDomainBlock, api.PostDomainBlock, scope.PostDomainBlock)
	register(api, e.DELETE, route.DeleteDomainBlock, api.DeleteDomainBlock, scope.DeleteDomainBlock)

	// https://docs.joinmastodon.org/methods/emails/
	register(api, e.POST, route.PostEmailConfirmation, api.PostEmailConfirmation, scope.PostEmailConfirmation)

	// https://docs.joinmastodon.org/methods/endorsements/
	register(api, e.GET, route.GetEndorsements, api.GetEndorsements, scope.GetEndorsements)

	// https://docs.joinmastodon.org/methods/favourites/
	register(api, e.GET, route.GetFavourites, api.GetFavourites, scope.GetFavourites)

	// https://docs.joinmastodon.org/methods/featured_tags/
	register(api, e.GET, route.GetFeaturedTags, api.GetFeaturedTags, scope.GetFeaturedTags)
	register(api, e.POST, route.PostFeaturedTag, api.PostFeaturedTag, scope.PostFeaturedTag)
	register(api, e.DELETE, route.DeleteFeaturedTag, api.DeleteFeaturedTag, scope.DeleteFeaturedTag)
	register(api, e.GET, route.GetFeaturedTags_Suggestions, api.GetFeaturedTags_Suggestions, scope.GetFeaturedTags_Suggestions)

	// https://docs.joinmastodon.org/methods/filters/
	register(api, e.GET, route.GetFilters, api.GetFilters, scope.GetFilters)
	register(api, e.GET, route.GetFilter, api.GetFilter, scope.GetFilter)
	register(api, e.POST, route.PostFilter, api.PostFilter, scope.PostFilter)
	register(api, e.PUT, route.PutFilter, api.PutFilter, scope.PutFilter)
	register(api, e.DELETE, route.DeleteFilter, api.DeleteFilter, scope.DeleteFilter)
	register(api, e.GET, route.GetFilter_Keywords, api.GetFilter_Keywords, scope.GetFilter_Keywords)
	register(api, e.POST, route.PostFilter_Keyword, api.PostFilter_Keyword, scope.PostFilter_Keyword)
	register(api, e.GET, route.GetFilter_Keyword, api.GetFilter_Keyword, scope.GetFilter_Keyword)
	register(api, e.PUT, route.PutFilter_Keyword, api.PutFilter_Keyword, scope.PutFilter_Keyword)
	register(api, e.DELETE, route.DeleteFilter_Keyword, api.DeleteFilter_Keyword, scope.DeleteFilter_Keyword)
	register(api, e.GET, route.GetFilter_Statuses, api.GetFilter_Statuses, scope.GetFilter_Statuses)
	register(api, e.POST, route.PostFilter_Status, api.PostFilter_Status, scope.PostFilter_Status)
	register(api, e.GET, route.GetFilter_Status, api.GetFilter_Status, scope.GetFilter_Status)
	register(api, e.DELETE, route.DeleteFilter_Status, api.DeleteFilter_Status, scope.DeleteFilter_Status)
	register(api, e.GET, route.GetFilter_V1, api.GetFilter_V1, scope.GetFilter_V1)
	register(api, e.POST, route.PostFilter_V1, api.PostFilter_V1, scope.PostFilter_V1)
	register(api, e.PUT, route.PutFilter_V1, api.PutFilter_V1, scope.PutFilter_V1)
	register(api, e.DELETE, route.DeleteFilter_V1, api.DeleteFilter_V1, scope.DeleteFilter)

	// https://docs.joinmastodon.org/methods/follow_requests/
	register(api, e.GET, route.GetFollowRequests, api.GetFollowRequests, scope.GetFollowRequests)
	register(api, e.POST, route.PostFollowRequest_Authorize, api.PostFollowRequest_Authorize, scope.PostFollowRequest_Authorize)
	register(api, e.POST, route.PostFollowRequest_Reject, api.PostFollowRequest_Reject, scope.PostFollowRequest_Reject)

	// https://docs.joinmastodon.org/methods/followed_tags/
	register(api, e.GET, route.GetFollowedTags, api.GetFollowedTags, scope.GetFollowedTags)

	// https://docs.joinmastodon.org/methods/instance/
	register(api, e.GET, route.GetInstance, api.GetInstance, scope.GetInstance)
	register(api, e.GET, route.GetInstance_Peers, api.GetInstance_Peers, scope.GetInstance_Peers)
	register(api, e.GET, route.GetInstance_Activity, api.GetInstance_Activity, scope.GetInstance_Activity)
	register(api, e.GET, route.GetInstance_Rules, api.GetInstance_Rules, scope.GetInstance_Rules)
	register(api, e.GET, route.GetInstance_DomainBlocks, api.GetInstance_DomainBlocks, scope.GetInstance_DomainBlocks)
	register(api, e.GET, route.GetInstance_ExtendedDescription, api.GetInstance_ExtendedDescription, scope.GetInstance_ExtendedDescription)

	// https://docs.joinmastodon.org/methods/lists/
	register(api, e.GET, route.GetLists, api.GetLists, scope.GetLists)
	register(api, e.GET, route.GetList, api.GetList, scope.GetList)
	register(api, e.POST, route.PostList, api.PostList, scope.PostList)
	register(api, e.PUT, route.PutList, api.PutList, scope.PutList)
	register(api, e.DELETE, route.DeleteList, api.DeleteList, scope.DeleteList)
	register(api, e.GET, route.GetList_Accounts, api.GetList_Accounts, scope.GetList_Accounts)
	register(api, e.POST, route.PostList_Accounts, api.PostList_Accounts, scope.PostList_Accounts)
	register(api, e.DELETE, route.DeleteList_Accounts, api.DeleteList_Accounts, scope.DeleteList_Accounts)

	// https://docs.joinmastodon.org/methods/markers/
	register(api, e.GET, route.GetMarkers, api.GetMarkers, scope.GetMarkers)
	register(api, e.POST, route.PostMarker, api.PostMarker, scope.PostMarker)

	// https://docs.joinmastodon.org/methods/media/
	register(api, e.POST, route.PostMedia, api.PostMedia, scope.PostMedia)

	// https://docs.joinmastodon.org/methods/mutes/
	register(api, e.GET, route.GetMutes, api.GetMutes, scope.GetMutes)

	// https://docs.joinmastodon.org/methods/notifications/
	register(api, e.GET, route.GetNotifications, api.GetNotifications, scope.GetNotifications)
	register(api, e.GET, route.GetNotification, api.GetNotification, scope.GetNotification)
	register(api, e.POST, route.PostNotifications_Clear, api.PostNotifications_Clear, scope.PostNotifications_Clear)
	register(api, e.POST, route.PostNotification_Dismiss, api.PostNotification_Dismiss, scope.PostNotification_Dismiss)

	// https://docs.joinmastodon.org/methods/oauth/
	register(api, e.GET, route.GetOAuth_Authorize, api.GetOAuth_Authorize, scope.GetOAuth_Authorize)
	register(api, e.POST, route.PostOAuth_Token, api.PostOAuth_Token, scope.PostOAuth_Token)
	register(api, e.POST, route.PostOAuth_Revoke, api.PostOAuth_Revoke, scope.PostOAuth_Revoke)

	// https://docs.joinmastodon.org/methods/oembed/
	register(api, e.GET, route.GetOEmbed, api.GetOEmbed, scope.GetOEmbed)

	// https://docs.joinmastodon.org/methods/polls/
	register(api, e.GET, route.GetPoll, api.GetPoll, scope.GetPoll)
	register(api, e.POST, route.PostPoll_Votes, api.PostPoll_Votes, scope.PostPoll_Votes)

	// https://docs.joinmastodon.org/methods/preferences/
	register(api, e.GET, route.GetPreferences, api.GetPreferences, scope.GetPreferences)

	// https://docs.joinmastodon.org/methods/profile/
	register(api, e.DELETE, route.DeleteProfile_Avatar, api.DeleteProfile_Avatar, scope.DeleteProfile_Avatar)
	register(api, e.DELETE, route.DeleteProfile_Header, api.DeleteProfile_Header, scope.DeleteProfile_Header)

	// https://docs.joinmastodon.org/methods/reports/
	register(api, e.POST, route.PostReport, api.PostReport, scope.PostReport)

	// https://docs.joinmastodon.org/methods/scheduled_statuses/
	register(api, e.GET, route.GetScheduledStatuses, api.GetScheduledStatuses, scope.GetScheduledStatuses)
	register(api, e.GET, route.GetScheduledStatus, api.GetScheduledStatus, scope.GetScheduledStatus)
	register(api, e.PUT, route.PutScheduledStatus, api.PutScheduledStatus, scope.PutScheduledStatus)
	register(api, e.DELETE, route.DeleteScheduledStatus, api.DeleteScheduledStatus, scope.DeleteScheduledStatus)

	// https://docs.joinmastodon.org/methods/search/
	register(api, e.GET, route.GetSearch, api.GetSearch, scope.GetSearch)

	// https://docs.joinmastodon.org/methods/statuses/#create
	register(api, e.POST, route.PostStatus, api.PostStatus, scope.PostStatus)
	register(api, e.GET, route.GetStatus, api.GetStatus, scope.GetStatus)
	register(api, e.DELETE, route.DeleteStatus, api.DeleteStatus, scope.DeleteStatus)
	register(api, e.GET, route.GetStatus_Context, api.GetStatus_Context, scope.GetStatus_Context)
	register(api, e.POST, route.PostStatus_Translate, api.PostStatus_Translate, scope.PostStatus_Translate)
	register(api, e.GET, route.GetStatus_RebloggedBy, api.GetStatus_RebloggedBy, scope.GetStatus_RebloggedBy)
	register(api, e.GET, route.GetStatus_FavouritedBy, api.GetStatus_FavouritedBy, scope.GetStatus_RebloggedBy)
	register(api, e.POST, route.PostStatus_Favourite, api.PostStatus_Favourite, scope.PostStatus_Favourite)
	register(api, e.POST, route.PostStatus_Unfavourite, api.PostStatus_Unfavourite, scope.PostStatus_Unfavourite)
	register(api, e.POST, route.PostStatus_Reblog, api.PostStatus_Reblog, scope.PostStatus_Reblog)
	register(api, e.POST, route.PostStatus_Unreblog, api.PostStatus_Unreblog, scope.PostStatus_Unreblog)
	register(api, e.POST, route.PostStatus_Bookmark, api.PostStatus_Bookmark, scope.PostStatus_Bookmark)
	register(api, e.POST, route.PostStatus_Unbookmark, api.PostStatus_Unbookmark, scope.PostStatus_Unbookmark)
	register(api, e.POST, route.PostStatus_Mute, api.PostStatus_Mute, scope.PostStatus_Mute)
	register(api, e.POST, route.PostStatus_Unmute, api.PostStatus_Unmute, scope.PostStatus_Unmute)
	register(api, e.POST, route.PostStatus_Pin, api.PostStatus_Pin, scope.PostStatus_Pin)
	register(api, e.POST, route.PostStatus_Unpin, api.PostStatus_Unpin, scope.PostStatus_Unpin)
	register(api, e.PUT, route.PutStatus, api.PutStatus, scope.PutStatus)
	register(api, e.GET, route.GetStatus_History, api.GetStatus_History, scope.GetStatus_History)
	register(api, e.GET, route.GetStatus_Source, api.GetStatus_Source, scope.GetStatus_Source)

	// https://docs.joinmastodon.org/methods/suggestions/
	register(api, e.GET, route.GetSuggestions, api.GetSuggestions, scope.GetSuggestions)
	register(api, e.DELETE, route.DeleteSuggestion, api.DeleteSuggestion, scope.DeleteSuggestion)

	// https://docs.joinmastodon.org/methods/tags/
	register(api, e.GET, route.GetTag, api.GetTag, scope.GetTag)
	register(api, e.POST, route.PostTag_Follow, api.PostTag_Follow, scope.PostTag_Follow)
	register(api, e.POST, route.PostTag_Unfollow, api.PostTag_Unfollow, scope.PostTag_Unfollow)

	// https://docs.joinmastodon.org/methods/timelines/
	register(api, e.GET, route.GetTimeline_Public, api.GetTimeline_Public, scope.GetTimeline_Public)
	register(api, e.GET, route.GetTimeline_Hashtag, api.GetTimeline_Hashtag, scope.GetTimeline_Hashtag)
	register(api, e.GET, route.GetTimeline_Home, api.GetTimeline_Home, scope.GetTimeline_Home)
	register(api, e.GET, route.GetTimeline_List, api.GetTimeline_List, scope.GetTimeline_List)

	// https://docs.joinmastodon.org/methods/trends/
	register(api, e.GET, route.GetTrends, api.GetTrends, scope.GetTrends)
	register(api, e.GET, route.GetTrends_Statuses, api.GetTrends_Statuses, scope.GetTrends_Statuses)
	register(api, e.GET, route.GetTrends_Links, api.GetTrends_Links, scope.GetTrends_Links)
}
