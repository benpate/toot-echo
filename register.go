package tootecho

import (
	"github.com/benpate/toot"
	"github.com/benpate/toot/route"
	"github.com/benpate/toot/scope"
	"github.com/labstack/echo/v4"
)

func Register[AuthToken toot.ScopesGetter](e *echo.Echo, api toot.API[AuthToken], middleware ...echo.MiddlewareFunc) {

	// https://docs.joinmastodon.org/methods/accounts/
	single_result(api, e.POST, route.PostAccount, api.PostAccount, scope.PostAccount)
	single_result(api, e.GET, route.GetAccount_VerifyCredentials, api.GetAccount_VerifyCredentials, scope.GetAccount_VerifyCredentials)
	single_result(api, e.PATCH, route.PatchAccount_UpdateCredentials, api.PatchAccount_UpdateCredentials, scope.PatchAccount_UpdateCredentials)
	single_result(api, e.GET, route.GetAccount, api.GetAccount, scope.GetAccount)
	paged_result(api, e.GET, route.GetAccount_Statuses, api.GetAccount_Statuses, scope.GetAccount_Statuses)
	paged_result(api, e.GET, route.GetAccount_Followers, api.GetAccount_Followers, scope.GetAccount_Followers)
	paged_result(api, e.GET, route.GetAccount_Following, api.GetAccount_Following, scope.GetAccount_Following)
	paged_result(api, e.GET, route.GetAccount_FeaturedTags, api.GetAccount_FeaturedTags, scope.GetAccount_FeaturedTags)
	single_result(api, e.POST, route.PostAccount, api.PostAccount_Follow, scope.PostAccont_Follow)
	single_result(api, e.POST, route.PostAccount_Unfollow, api.PostAccount_Unfollow, scope.PostAccount_Unfollow)
	single_result(api, e.POST, route.PostAccount_Block, api.PostAccount_Block, scope.PostAccount_Block)
	single_result(api, e.POST, route.PostAccount_Unblock, api.PostAccount_Unblock, scope.PostAccount_Unblock)
	single_result(api, e.POST, route.PostAccount_Mute, api.PostAccount_Mute, scope.PostAccount_Mute)
	single_result(api, e.POST, route.PostAccount_Unmute, api.PostAccount_Unmute, scope.PostAccount_Unmute)
	single_result(api, e.POST, route.PostAccount_Pin, api.PostAccount_Pin, scope.PostAccount_Pin)
	single_result(api, e.POST, route.PostAccount_Unpin, api.PostAccount_Unpin, scope.PostAccount_Unpin)
	single_result(api, e.POST, route.PostAccount_Note, api.PostAccount_Note, scope.PostAccount_Note)
	single_result(api, e.GET, route.PostAccount_Relationships, api.GetAccount_Relationships, scope.PostAccount_Relationships)
	single_result(api, e.GET, route.GetAccount_FamiliarFollowers, api.GetAccount_FamiliarFollowers, scope.GetAccount_FamiliarFollowers)
	paged_result(api, e.GET, route.GetAccount_Search, api.GetAccount_Search, scope.GetAccount_Search)
	single_result(api, e.GET, route.GetAccount_Lookup, api.GetAccount_Lookup, scope.GetAccount_Lookup)

	// https://docs.joinmastodon.org/methods/apps/
	single_result(api, e.POST, route.PostApplication, api.PostApplication, scope.PostApplication)
	single_result(api, e.GET, route.GetApplication_VerifyCredentials, api.GetApplication_VerifyCredentials, scope.GetApplication_VerifyCredentials)

	// https://docs.joinmastodon.org/methods/announcements/
	single_result(api, e.GET, route.GetAnnouncements, api.GetAnnouncements, scope.GetAnnouncements)
	single_result(api, e.POST, route.PostAnnouncement_Dismiss, api.PostAnnouncement_Dismiss, scope.PostAnnoucement_Dismis)
	single_result(api, e.PUT, route.PutAnnouncement_Reaction, api.PutAnnouncement_Reaction, scope.PutAnnouncement_Reaction)
	single_result(api, e.DELETE, route.DeleteAnnouncement_Reaction, api.DeleteAnnouncement_Reaction, scope.DeleteAnnouncement_Reaction)

	// https://docs.joinmastodon.org/methods/apps/
	single_result(api, e.POST, route.PostApplication, api.PostApplication, scope.PostApplication)
	single_result(api, e.GET, route.GetApplication_VerifyCredentials, api.GetApplication_VerifyCredentials, scope.GetApplication_VerifyCredentials)

	// https://docs.joinmastodon.org/methods/blocks/
	paged_result(api, e.GET, route.GetBlocks, api.GetBlocks, scope.GetBlocks)

	// https://docs.joinmastodon.org/methods/bookmarks/
	single_result(api, e.GET, route.GetBookmarks, api.GetBookmarks, scope.GetBookmarks)

	// https://docs.joinmastodon.org/methods/conversations/
	paged_result(api, e.GET, route.GetConversations, api.GetConversations, scope.GetConversations)
	single_result(api, e.DELETE, route.DeleteConversation, api.DeleteConversation, scope.DeleteConversation)
	single_result(api, e.POST, route.PostConversationRead, api.PostConversationRead, scope.PostConversationRead)

	// https://docs.joinmastodon.org/methods/custom_emojis/
	single_result(api, e.GET, route.GetCustomEmojis, api.GetCustomEmojis, scope.GetCustomEmojis)

	// https://docs.joinmastodon.org/methods/directory/
	paged_result(api, e.GET, route.GetDirectory, api.GetDirectory, scope.GetDirectory)

	// https://docs.joinmastodon.org/methods/domain_blocks/
	paged_result(api, e.GET, route.GetDomainBlocks, api.GetDomainBlocks, scope.GetDomainBlocks)
	single_result(api, e.POST, route.PostDomainBlock, api.PostDomainBlock, scope.PostDomainBlock)
	single_result(api, e.DELETE, route.DeleteDomainBlock, api.DeleteDomainBlock, scope.DeleteDomainBlock)

	// https://docs.joinmastodon.org/methods/emails/
	single_result(api, e.POST, route.PostEmailConfirmation, api.PostEmailConfirmation, scope.PostEmailConfirmation)

	// https://docs.joinmastodon.org/methods/endorsements/
	paged_result(api, e.GET, route.GetEndorsements, api.GetEndorsements, scope.GetEndorsements)

	// https://docs.joinmastodon.org/methods/favourites/
	single_result(api, e.GET, route.GetFavourites, api.GetFavourites, scope.GetFavourites)

	// https://docs.joinmastodon.org/methods/featured_tags/
	single_result(api, e.GET, route.GetFeaturedTags, api.GetFeaturedTags, scope.GetFeaturedTags)
	single_result(api, e.POST, route.PostFeaturedTag, api.PostFeaturedTag, scope.PostFeaturedTag)
	single_result(api, e.DELETE, route.DeleteFeaturedTag, api.DeleteFeaturedTag, scope.DeleteFeaturedTag)
	single_result(api, e.GET, route.GetFeaturedTags_Suggestions, api.GetFeaturedTags_Suggestions, scope.GetFeaturedTags_Suggestions)

	// https://docs.joinmastodon.org/methods/filters/
	single_result(api, e.GET, route.GetFilters, api.GetFilters, scope.GetFilters)
	single_result(api, e.GET, route.GetFilter, api.GetFilter, scope.GetFilter)
	single_result(api, e.POST, route.PostFilter, api.PostFilter, scope.PostFilter)
	single_result(api, e.PUT, route.PutFilter, api.PutFilter, scope.PutFilter)
	single_result(api, e.DELETE, route.DeleteFilter, api.DeleteFilter, scope.DeleteFilter)
	single_result(api, e.GET, route.GetFilter_Keywords, api.GetFilter_Keywords, scope.GetFilter_Keywords)
	single_result(api, e.POST, route.PostFilter_Keyword, api.PostFilter_Keyword, scope.PostFilter_Keyword)
	single_result(api, e.GET, route.GetFilter_Keyword, api.GetFilter_Keyword, scope.GetFilter_Keyword)
	single_result(api, e.PUT, route.PutFilter_Keyword, api.PutFilter_Keyword, scope.PutFilter_Keyword)
	single_result(api, e.DELETE, route.DeleteFilter_Keyword, api.DeleteFilter_Keyword, scope.DeleteFilter_Keyword)
	single_result(api, e.GET, route.GetFilter_Statuses, api.GetFilter_Statuses, scope.GetFilter_Statuses)
	single_result(api, e.POST, route.PostFilter_Status, api.PostFilter_Status, scope.PostFilter_Status)
	single_result(api, e.GET, route.GetFilter_Status, api.GetFilter_Status, scope.GetFilter_Status)
	single_result(api, e.DELETE, route.DeleteFilter_Status, api.DeleteFilter_Status, scope.DeleteFilter_Status)
	single_result(api, e.GET, route.GetFilter_V1, api.GetFilter_V1, scope.GetFilter_V1)
	single_result(api, e.POST, route.PostFilter_V1, api.PostFilter_V1, scope.PostFilter_V1)
	single_result(api, e.PUT, route.PutFilter_V1, api.PutFilter_V1, scope.PutFilter_V1)
	single_result(api, e.DELETE, route.DeleteFilter_V1, api.DeleteFilter_V1, scope.DeleteFilter)

	// https://docs.joinmastodon.org/methods/follow_requests/
	paged_result(api, e.GET, route.GetFollowRequests, api.GetFollowRequests, scope.GetFollowRequests)
	single_result(api, e.POST, route.PostFollowRequest_Authorize, api.PostFollowRequest_Authorize, scope.PostFollowRequest_Authorize)
	single_result(api, e.POST, route.PostFollowRequest_Reject, api.PostFollowRequest_Reject, scope.PostFollowRequest_Reject)

	// https://docs.joinmastodon.org/methods/followed_tags/
	paged_result(api, e.GET, route.GetFollowedTags, api.GetFollowedTags, scope.GetFollowedTags)

	// https://docs.joinmastodon.org/methods/instance/
	single_result(api, e.GET, route.GetInstance, api.GetInstance, scope.GetInstance)
	single_result(api, e.GET, route.GetInstance_Peers, api.GetInstance_Peers, scope.GetInstance_Peers)
	single_result(api, e.GET, route.GetInstance_Activity, api.GetInstance_Activity, scope.GetInstance_Activity)
	single_result(api, e.GET, route.GetInstance_Rules, api.GetInstance_Rules, scope.GetInstance_Rules)
	single_result(api, e.GET, route.GetInstance_DomainBlocks, api.GetInstance_DomainBlocks, scope.GetInstance_DomainBlocks)
	single_result(api, e.GET, route.GetInstance_ExtendedDescription, api.GetInstance_ExtendedDescription, scope.GetInstance_ExtendedDescription)
	single_result(api, e.GET, route.GetInstance_V1, api.GetInstance_V1, scope.GetInstance_V1)

	// https://docs.joinmastodon.org/methods/lists/
	single_result(api, e.GET, route.GetLists, api.GetLists, scope.GetLists)
	single_result(api, e.GET, route.GetList, api.GetList, scope.GetList)
	single_result(api, e.POST, route.PostList, api.PostList, scope.PostList)
	single_result(api, e.PUT, route.PutList, api.PutList, scope.PutList)
	single_result(api, e.DELETE, route.DeleteList, api.DeleteList, scope.DeleteList)
	paged_result(api, e.GET, route.GetList_Accounts, api.GetList_Accounts, scope.GetList_Accounts)
	single_result(api, e.POST, route.PostList_Accounts, api.PostList_Accounts, scope.PostList_Accounts)
	single_result(api, e.DELETE, route.DeleteList_Accounts, api.DeleteList_Accounts, scope.DeleteList_Accounts)

	// https://docs.joinmastodon.org/methods/markers/
	single_result(api, e.GET, route.GetMarkers, api.GetMarkers, scope.GetMarkers)
	single_result(api, e.POST, route.PostMarker, api.PostMarker, scope.PostMarker)

	// https://docs.joinmastodon.org/methods/media/
	single_result(api, e.POST, route.PostMedia, api.PostMedia, scope.PostMedia)

	// https://docs.joinmastodon.org/methods/mutes/
	paged_result(api, e.GET, route.GetMutes, api.GetMutes, scope.GetMutes)

	// https://docs.joinmastodon.org/methods/notifications/
	paged_result(api, e.GET, route.GetNotifications, api.GetNotifications, scope.GetNotifications)
	single_result(api, e.GET, route.GetNotification, api.GetNotification, scope.GetNotification)
	single_result(api, e.POST, route.PostNotifications_Clear, api.PostNotifications_Clear, scope.PostNotifications_Clear)
	single_result(api, e.POST, route.PostNotification_Dismiss, api.PostNotification_Dismiss, scope.PostNotification_Dismiss)

	// https://docs.joinmastodon.org/methods/oauth/
	single_result(api, e.GET, route.GetOAuth_Authorize, api.GetOAuth_Authorize, scope.GetOAuth_Authorize)
	single_result(api, e.POST, route.PostOAuth_Token, api.PostOAuth_Token, scope.PostOAuth_Token)
	single_result(api, e.POST, route.PostOAuth_Revoke, api.PostOAuth_Revoke, scope.PostOAuth_Revoke)

	// https://docs.joinmastodon.org/methods/oembed/
	single_result(api, e.GET, route.GetOEmbed, api.GetOEmbed, scope.GetOEmbed)

	// https://docs.joinmastodon.org/methods/polls/
	single_result(api, e.GET, route.GetPoll, api.GetPoll, scope.GetPoll)
	single_result(api, e.POST, route.PostPoll_Votes, api.PostPoll_Votes, scope.PostPoll_Votes)

	// https://docs.joinmastodon.org/methods/preferences/
	single_result(api, e.GET, route.GetPreferences, api.GetPreferences, scope.GetPreferences)

	// https://docs.joinmastodon.org/methods/profile/
	single_result(api, e.DELETE, route.DeleteProfile_Avatar, api.DeleteProfile_Avatar, scope.DeleteProfile_Avatar)
	single_result(api, e.DELETE, route.DeleteProfile_Header, api.DeleteProfile_Header, scope.DeleteProfile_Header)

	// https://docs.joinmastodon.org/methods/reports/
	single_result(api, e.POST, route.PostReport, api.PostReport, scope.PostReport)

	// https://docs.joinmastodon.org/methods/scheduled_statuses/
	paged_result(api, e.GET, route.GetScheduledStatuses, api.GetScheduledStatuses, scope.GetScheduledStatuses)
	single_result(api, e.GET, route.GetScheduledStatus, api.GetScheduledStatus, scope.GetScheduledStatus)
	single_result(api, e.PUT, route.PutScheduledStatus, api.PutScheduledStatus, scope.PutScheduledStatus)
	single_result(api, e.DELETE, route.DeleteScheduledStatus, api.DeleteScheduledStatus, scope.DeleteScheduledStatus)

	// https://docs.joinmastodon.org/methods/search/
	single_result(api, e.GET, route.GetSearch, api.GetSearch, scope.GetSearch)

	// https://docs.joinmastodon.org/methods/statuses/
	single_result(api, e.POST, route.PostStatus, api.PostStatus, scope.PostStatus)
	single_result(api, e.GET, route.GetStatus, api.GetStatus, scope.GetStatus)
	single_result(api, e.DELETE, route.DeleteStatus, api.DeleteStatus, scope.DeleteStatus)
	single_result(api, e.GET, route.GetStatus_Context, api.GetStatus_Context, scope.GetStatus_Context)
	single_result(api, e.POST, route.PostStatus_Translate, api.PostStatus_Translate, scope.PostStatus_Translate)
	paged_result(api, e.GET, route.GetStatus_RebloggedBy, api.GetStatus_RebloggedBy, scope.GetStatus_RebloggedBy)
	paged_result(api, e.GET, route.GetStatus_FavouritedBy, api.GetStatus_FavouritedBy, scope.GetStatus_RebloggedBy)
	single_result(api, e.POST, route.PostStatus_Favourite, api.PostStatus_Favourite, scope.PostStatus_Favourite)
	single_result(api, e.POST, route.PostStatus_Unfavourite, api.PostStatus_Unfavourite, scope.PostStatus_Unfavourite)
	single_result(api, e.POST, route.PostStatus_Reblog, api.PostStatus_Reblog, scope.PostStatus_Reblog)
	single_result(api, e.POST, route.PostStatus_Unreblog, api.PostStatus_Unreblog, scope.PostStatus_Unreblog)
	single_result(api, e.POST, route.PostStatus_Bookmark, api.PostStatus_Bookmark, scope.PostStatus_Bookmark)
	single_result(api, e.POST, route.PostStatus_Unbookmark, api.PostStatus_Unbookmark, scope.PostStatus_Unbookmark)
	single_result(api, e.POST, route.PostStatus_Mute, api.PostStatus_Mute, scope.PostStatus_Mute)
	single_result(api, e.POST, route.PostStatus_Unmute, api.PostStatus_Unmute, scope.PostStatus_Unmute)
	single_result(api, e.POST, route.PostStatus_Pin, api.PostStatus_Pin, scope.PostStatus_Pin)
	single_result(api, e.POST, route.PostStatus_Unpin, api.PostStatus_Unpin, scope.PostStatus_Unpin)
	single_result(api, e.PUT, route.PutStatus, api.PutStatus, scope.PutStatus)
	single_result(api, e.GET, route.GetStatus_History, api.GetStatus_History, scope.GetStatus_History)
	single_result(api, e.GET, route.GetStatus_Source, api.GetStatus_Source, scope.GetStatus_Source)

	// https://docs.joinmastodon.org/methods/suggestions/
	single_result(api, e.GET, route.GetSuggestions, api.GetSuggestions, scope.GetSuggestions)
	single_result(api, e.DELETE, route.DeleteSuggestion, api.DeleteSuggestion, scope.DeleteSuggestion)

	// https://docs.joinmastodon.org/methods/tags/
	single_result(api, e.GET, route.GetTag, api.GetTag, scope.GetTag)
	single_result(api, e.POST, route.PostTag_Follow, api.PostTag_Follow, scope.PostTag_Follow)
	single_result(api, e.POST, route.PostTag_Unfollow, api.PostTag_Unfollow, scope.PostTag_Unfollow)

	// https://docs.joinmastodon.org/methods/timelines/
	paged_result(api, e.GET, route.GetTimeline_Public, api.GetTimeline_Public, scope.GetTimeline_Public)
	paged_result(api, e.GET, route.GetTimeline_Hashtag, api.GetTimeline_Hashtag, scope.GetTimeline_Hashtag)
	paged_result(api, e.GET, route.GetTimeline_Home, api.GetTimeline_Home, scope.GetTimeline_Home)
	paged_result(api, e.GET, route.GetTimeline_List, api.GetTimeline_List, scope.GetTimeline_List)

	// https://docs.joinmastodon.org/methods/trends/
	paged_result(api, e.GET, route.GetTrends, api.GetTrends, scope.GetTrends)
	paged_result(api, e.GET, route.GetTrends_Statuses, api.GetTrends_Statuses, scope.GetTrends_Statuses)
	paged_result(api, e.GET, route.GetTrends_Links, api.GetTrends_Links, scope.GetTrends_Links)
}
