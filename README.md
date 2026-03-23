# Zernio Go SDK

Official Go client library for the [Zernio API](https://docs.zernio.com) - Schedule and manage social media posts across multiple platforms.

## Installation

```bash
go get github.com/zernio-dev/zernio-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    zernio "github.com/zernio-dev/zernio-go/zernio"
)

func main() {
    client, err := zernio.NewClientWithResponses("https://api.zernio.com",
        zernio.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
            req.Header.Set("Authorization", "Bearer YOUR_API_KEY")
            return nil
        }),
    )
    if err != nil {
        log.Fatal(err)
    }

    // List accounts
    resp, err := client.ListAccountsWithResponse(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Accounts: %+v\n", resp.JSON200)
}
```

## SDK Reference

### Posts
| Method | Description |
|--------|-------------|
| `client.ListPostsWithResponse()` | List posts |
| `client.BulkUploadPostsWithResponse()` | Bulk upload from CSV |
| `client.CreatePostWithResponse()` | Create post |
| `client.GetPostWithResponse()` | Get post |
| `client.UpdatePostWithResponse()` | Update post |
| `client.DeletePostWithResponse()` | Delete post |
| `client.RetryPostWithResponse()` | Retry failed post |
| `client.UnpublishPostWithResponse()` | Unpublish post |

### Accounts
| Method | Description |
|--------|-------------|
| `client.GetAllAccountsHealthWithResponse()` | Check accounts health |
| `client.ListAccountsWithResponse()` | List accounts |
| `client.GetAccountHealthWithResponse()` | Check account health |
| `client.GetFollowerStatsWithResponse()` | Get follower stats |
| `client.GetGoogleBusinessReviewsWithResponse()` | Get reviews |
| `client.GetLinkedInMentionsWithResponse()` | Resolve LinkedIn mention |
| `client.GetTikTokCreatorInfoWithResponse()` | Get TikTok creator info |
| `client.UpdateAccountWithResponse()` | Update account |
| `client.DeleteAccountWithResponse()` | Disconnect account |

### Profiles
| Method | Description |
|--------|-------------|
| `client.ListProfilesWithResponse()` | List profiles |
| `client.CreateProfileWithResponse()` | Create profile |
| `client.GetProfileWithResponse()` | Get profile |
| `client.UpdateProfileWithResponse()` | Update profile |
| `client.DeleteProfileWithResponse()` | Delete profile |

### Analytics
| Method | Description |
|--------|-------------|
| `client.GetAnalyticsWithResponse()` | Get post analytics |
| `client.GetBestTimeToPostWithResponse()` | Get best times to post |
| `client.GetContentDecayWithResponse()` | Get content performance decay |
| `client.GetDailyMetricsWithResponse()` | Get daily aggregated metrics |
| `client.GetInstagramAccountInsightsWithResponse()` | Get Instagram account-level insights |
| `client.GetInstagramDemographicsWithResponse()` | Get Instagram audience demographics |
| `client.GetLinkedInAggregateAnalyticsWithResponse()` | Get LinkedIn aggregate stats |
| `client.GetLinkedInPostAnalyticsWithResponse()` | Get LinkedIn post stats |
| `client.GetLinkedInPostReactionsWithResponse()` | Get LinkedIn post reactions |
| `client.GetPostTimelineWithResponse()` | Get post analytics timeline |
| `client.GetPostingFrequencyWithResponse()` | Get posting frequency vs engagement |
| `client.GetYouTubeDailyViewsWithResponse()` | Get YouTube daily views |

### Account Groups
| Method | Description |
|--------|-------------|
| `client.ListAccountGroupsWithResponse()` | List groups |
| `client.CreateAccountGroupWithResponse()` | Create group |
| `client.UpdateAccountGroupWithResponse()` | Update group |
| `client.DeleteAccountGroupWithResponse()` | Delete group |

### Queue
| Method | Description |
|--------|-------------|
| `client.ListQueueSlotsWithResponse()` | List schedules |
| `client.CreateQueueSlotWithResponse()` | Create schedule |
| `client.GetNextQueueSlotWithResponse()` | Get next available slot |
| `client.UpdateQueueSlotWithResponse()` | Update schedule |
| `client.DeleteQueueSlotWithResponse()` | Delete schedule |
| `client.PreviewQueueWithResponse()` | Preview upcoming slots |

### Webhooks
| Method | Description |
|--------|-------------|
| `client.CreateWebhookSettingsWithResponse()` | Create webhook |
| `client.GetWebhookLogsWithResponse()` | Get delivery logs |
| `client.GetWebhookSettingsWithResponse()` | List webhooks |
| `client.UpdateWebhookSettingsWithResponse()` | Update webhook |
| `client.DeleteWebhookSettingsWithResponse()` | Delete webhook |
| `client.TestWebhookWithResponse()` | Send test webhook |

### API Keys
| Method | Description |
|--------|-------------|
| `client.ListApiKeysWithResponse()` | List keys |
| `client.CreateApiKeyWithResponse()` | Create key |
| `client.DeleteApiKeyWithResponse()` | Delete key |

### Media
| Method | Description |
|--------|-------------|
| `client.GetMediaPresignedUrlWithResponse()` | Get presigned upload URL |

### Tools
| Method | Description |
|--------|-------------|
| `client.GetYouTubeTranscriptWithResponse()` | Get YouTube transcript |
| `client.CheckInstagramHashtagsWithResponse()` | Check IG hashtag bans |
| `client.DownloadBlueskyMediaWithResponse()` | Download Bluesky media |
| `client.DownloadFacebookVideoWithResponse()` | Download Facebook video |
| `client.DownloadInstagramMediaWithResponse()` | Download Instagram media |
| `client.DownloadLinkedInVideoWithResponse()` | Download LinkedIn video |
| `client.DownloadTikTokVideoWithResponse()` | Download TikTok video |
| `client.DownloadTwitterMediaWithResponse()` | Download Twitter/X media |
| `client.DownloadYouTubeVideoWithResponse()` | Download YouTube video |

### Users
| Method | Description |
|--------|-------------|
| `client.ListUsersWithResponse()` | List users |
| `client.GetUserWithResponse()` | Get user |

### Usage
| Method | Description |
|--------|-------------|
| `client.GetUsageStatsWithResponse()` | Get plan and usage stats |

### Logs
| Method | Description |
|--------|-------------|
| `client.ListConnectionLogsWithResponse()` | List connection logs |
| `client.ListPostsLogsWithResponse()` | List publishing logs |
| `client.GetPostLogsWithResponse()` | Get post logs |

### Connect (OAuth)
| Method | Description |
|--------|-------------|
| `client.ListFacebookPagesWithResponse()` | List Facebook pages |
| `client.ListGoogleBusinessLocationsWithResponse()` | List GBP locations |
| `client.ListLinkedInOrganizationsWithResponse()` | List LinkedIn orgs |
| `client.ListPinterestBoardsForSelectionWithResponse()` | List Pinterest boards |
| `client.ListSnapchatProfilesWithResponse()` | List Snapchat profiles |
| `client.GetConnectUrlWithResponse()` | Get OAuth connect URL |
| `client.GetFacebookPagesWithResponse()` | List Facebook pages |
| `client.GetGmbLocationsWithResponse()` | List GBP locations |
| `client.GetLinkedInOrganizationsWithResponse()` | List LinkedIn orgs |
| `client.GetPendingOAuthDataWithResponse()` | Get pending OAuth data |
| `client.GetPinterestBoardsWithResponse()` | List Pinterest boards |
| `client.GetRedditFlairsWithResponse()` | List subreddit flairs |
| `client.GetRedditSubredditsWithResponse()` | List Reddit subreddits |
| `client.GetTelegramConnectStatusWithResponse()` | Generate Telegram code |
| `client.UpdateFacebookPageWithResponse()` | Update Facebook page |
| `client.UpdateGmbLocationWithResponse()` | Update GBP location |
| `client.UpdateLinkedInOrganizationWithResponse()` | Switch LinkedIn account type |
| `client.UpdatePinterestBoardsWithResponse()` | Set default Pinterest board |
| `client.UpdateRedditSubredditsWithResponse()` | Set default subreddit |
| `client.CompleteTelegramConnectWithResponse()` | Check Telegram status |
| `client.ConnectBlueskyCredentialsWithResponse()` | Connect Bluesky account |
| `client.ConnectWhatsAppCredentialsWithResponse()` | Connect WhatsApp via credentials |
| `client.HandleOAuthCallbackWithResponse()` | Complete OAuth callback |
| `client.InitiateTelegramConnectWithResponse()` | Connect Telegram directly |
| `client.SelectFacebookPageWithResponse()` | Select Facebook page |
| `client.SelectGoogleBusinessLocationWithResponse()` | Select GBP location |
| `client.SelectLinkedInOrganizationWithResponse()` | Select LinkedIn org |
| `client.SelectPinterestBoardWithResponse()` | Select Pinterest board |
| `client.SelectSnapchatProfileWithResponse()` | Select Snapchat profile |

### Reddit
| Method | Description |
|--------|-------------|
| `client.GetRedditFeedWithResponse()` | Get subreddit feed |
| `client.SearchRedditWithResponse()` | Search posts |

### Account Settings
| Method | Description |
|--------|-------------|
| `client.GetInstagramIceBreakersWithResponse()` | Get IG ice breakers |
| `client.GetMessengerMenuWithResponse()` | Get FB persistent menu |
| `client.GetTelegramCommandsWithResponse()` | Get TG bot commands |
| `client.DeleteInstagramIceBreakersWithResponse()` | Delete IG ice breakers |
| `client.DeleteMessengerMenuWithResponse()` | Delete FB persistent menu |
| `client.DeleteTelegramCommandsWithResponse()` | Delete TG bot commands |
| `client.SetInstagramIceBreakersWithResponse()` | Set IG ice breakers |
| `client.SetMessengerMenuWithResponse()` | Set FB persistent menu |
| `client.SetTelegramCommandsWithResponse()` | Set TG bot commands |

### Comments (Inbox)
| Method | Description |
|--------|-------------|
| `client.ListInboxCommentsWithResponse()` | List commented posts |
| `client.GetInboxPostCommentsWithResponse()` | Get post comments |
| `client.DeleteInboxCommentWithResponse()` | Delete comment |
| `client.HideInboxCommentWithResponse()` | Hide comment |
| `client.LikeInboxCommentWithResponse()` | Like comment |
| `client.ReplyToInboxPostWithResponse()` | Reply to comment |
| `client.SendPrivateReplyToCommentWithResponse()` | Send private reply |
| `client.UnhideInboxCommentWithResponse()` | Unhide comment |
| `client.UnlikeInboxCommentWithResponse()` | Unlike comment |

### GMB Attributes
| Method | Description |
|--------|-------------|
| `client.GetGoogleBusinessAttributesWithResponse()` | Get attributes |
| `client.UpdateGoogleBusinessAttributesWithResponse()` | Update attributes |

### GMB Food Menus
| Method | Description |
|--------|-------------|
| `client.GetGoogleBusinessFoodMenusWithResponse()` | Get food menus |
| `client.UpdateGoogleBusinessFoodMenusWithResponse()` | Update food menus |

### GMB Location Details
| Method | Description |
|--------|-------------|
| `client.GetGoogleBusinessLocationDetailsWithResponse()` | Get location details |
| `client.UpdateGoogleBusinessLocationDetailsWithResponse()` | Update location details |

### GMB Media
| Method | Description |
|--------|-------------|
| `client.ListGoogleBusinessMediaWithResponse()` | List media |
| `client.CreateGoogleBusinessMediaWithResponse()` | Upload photo |
| `client.DeleteGoogleBusinessMediaWithResponse()` | Delete photo |

### GMB Place Actions
| Method | Description |
|--------|-------------|
| `client.ListGoogleBusinessPlaceActionsWithResponse()` | List action links |
| `client.CreateGoogleBusinessPlaceActionWithResponse()` | Create action link |
| `client.DeleteGoogleBusinessPlaceActionWithResponse()` | Delete action link |

### Messages (Inbox)
| Method | Description |
|--------|-------------|
| `client.ListInboxConversationsWithResponse()` | List conversations |
| `client.GetInboxConversationWithResponse()` | Get conversation |
| `client.GetInboxConversationMessagesWithResponse()` | List messages |
| `client.UpdateInboxConversationWithResponse()` | Update conversation status |
| `client.EditInboxMessageWithResponse()` | Edit message |
| `client.SendInboxMessageWithResponse()` | Send message |

### Reviews (Inbox)
| Method | Description |
|--------|-------------|
| `client.ListInboxReviewsWithResponse()` | List reviews |
| `client.DeleteInboxReviewReplyWithResponse()` | Delete review reply |
| `client.ReplyToInboxReviewWithResponse()` | Reply to review |

### Twitter Engagement
| Method | Description |
|--------|-------------|
| `client.BookmarkPostWithResponse()` | Bookmark a tweet |
| `client.FollowUserWithResponse()` | Follow a user |
| `client.RemoveBookmarkWithResponse()` | Remove bookmark |
| `client.RetweetPostWithResponse()` | Retweet a post |
| `client.UndoRetweetWithResponse()` | Undo retweet |
| `client.UnfollowUserWithResponse()` | Unfollow a user |

### Validate
| Method | Description |
|--------|-------------|
| `client.ValidateMediaWithResponse()` | Validate media URL |
| `client.ValidatePostWithResponse()` | Validate post content |
| `client.ValidatePostLengthWithResponse()` | Validate post character count |
| `client.ValidateSubredditWithResponse()` | Check subreddit existence |

### WhatsApp
| Method | Description |
|--------|-------------|
| `client.BulkDeleteWhatsAppContactsWithResponse()` | Bulk delete contacts |
| `client.BulkUpdateWhatsAppContactsWithResponse()` | Bulk update contacts |
| `client.CreateWhatsAppBroadcastWithResponse()` | Create broadcast |
| `client.CreateWhatsAppContactWithResponse()` | Create contact |
| `client.CreateWhatsAppTemplateWithResponse()` | Create template |
| `client.GetWhatsAppBroadcastWithResponse()` | Get broadcast |
| `client.GetWhatsAppBroadcastRecipientsWithResponse()` | List recipients |
| `client.GetWhatsAppBroadcastsWithResponse()` | List broadcasts |
| `client.GetWhatsAppBusinessProfileWithResponse()` | Get business profile |
| `client.GetWhatsAppContactWithResponse()` | Get contact |
| `client.GetWhatsAppContactsWithResponse()` | List contacts |
| `client.GetWhatsAppDisplayNameWithResponse()` | Get display name and review status |
| `client.GetWhatsAppGroupsWithResponse()` | List contact groups |
| `client.GetWhatsAppTemplateWithResponse()` | Get template |
| `client.GetWhatsAppTemplatesWithResponse()` | List templates |
| `client.UpdateWhatsAppBusinessProfileWithResponse()` | Update business profile |
| `client.UpdateWhatsAppContactWithResponse()` | Update contact |
| `client.UpdateWhatsAppDisplayNameWithResponse()` | Request display name change |
| `client.UpdateWhatsAppTemplateWithResponse()` | Update template |
| `client.DeleteWhatsAppBroadcastWithResponse()` | Delete broadcast |
| `client.DeleteWhatsAppContactWithResponse()` | Delete contact |
| `client.DeleteWhatsAppGroupWithResponse()` | Delete group |
| `client.DeleteWhatsAppTemplateWithResponse()` | Delete template |
| `client.AddWhatsAppBroadcastRecipientsWithResponse()` | Add recipients |
| `client.CancelWhatsAppBroadcastScheduleWithResponse()` | Cancel scheduled broadcast |
| `client.ImportWhatsAppContactsWithResponse()` | Bulk import contacts |
| `client.RemoveWhatsAppBroadcastRecipientsWithResponse()` | Remove recipients |
| `client.RenameWhatsAppGroupWithResponse()` | Rename group |
| `client.ScheduleWhatsAppBroadcastWithResponse()` | Schedule broadcast |
| `client.SendWhatsAppBroadcastWithResponse()` | Send broadcast |
| `client.SendWhatsAppBulkWithResponse()` | Bulk send template messages |
| `client.UploadWhatsAppProfilePhotoWithResponse()` | Upload profile picture |

### WhatsApp Phone Numbers
| Method | Description |
|--------|-------------|
| `client.GetWhatsAppPhoneNumberWithResponse()` | Get phone number |
| `client.GetWhatsAppPhoneNumbersWithResponse()` | List phone numbers |
| `client.PurchaseWhatsAppPhoneNumberWithResponse()` | Purchase phone number |
| `client.ReleaseWhatsAppPhoneNumberWithResponse()` | Release phone number |

### Invites
| Method | Description |
|--------|-------------|
| `client.CreateInviteTokenWithResponse()` | Create invite token |

## SDK Reference

### Posts
| Method | Description |
|--------|-------------|
| `client.ListPostsWithResponse()` | List posts |
| `client.BulkUploadPostsWithResponse()` | Bulk upload from CSV |
| `client.CreatePostWithResponse()` | Create post |
| `client.GetPostWithResponse()` | Get post |
| `client.UpdatePostWithResponse()` | Update post |
| `client.DeletePostWithResponse()` | Delete post |
| `client.RetryPostWithResponse()` | Retry failed post |
| `client.UnpublishPostWithResponse()` | Unpublish post |

### Accounts
| Method | Description |
|--------|-------------|
| `client.GetAllAccountsHealthWithResponse()` | Check accounts health |
| `client.ListAccountsWithResponse()` | List accounts |
| `client.GetAccountHealthWithResponse()` | Check account health |
| `client.GetFollowerStatsWithResponse()` | Get follower stats |
| `client.GetGoogleBusinessReviewsWithResponse()` | Get reviews |
| `client.GetLinkedInMentionsWithResponse()` | Resolve LinkedIn mention |
| `client.GetTikTokCreatorInfoWithResponse()` | Get TikTok creator info |
| `client.UpdateAccountWithResponse()` | Update account |
| `client.DeleteAccountWithResponse()` | Disconnect account |

### Profiles
| Method | Description |
|--------|-------------|
| `client.ListProfilesWithResponse()` | List profiles |
| `client.CreateProfileWithResponse()` | Create profile |
| `client.GetProfileWithResponse()` | Get profile |
| `client.UpdateProfileWithResponse()` | Update profile |
| `client.DeleteProfileWithResponse()` | Delete profile |

### Analytics
| Method | Description |
|--------|-------------|
| `client.GetAnalyticsWithResponse()` | Get post analytics |
| `client.GetBestTimeToPostWithResponse()` | Get best times to post |
| `client.GetContentDecayWithResponse()` | Get content performance decay |
| `client.GetDailyMetricsWithResponse()` | Get daily aggregated metrics |
| `client.GetInstagramAccountInsightsWithResponse()` | Get Instagram account-level insights |
| `client.GetInstagramDemographicsWithResponse()` | Get Instagram audience demographics |
| `client.GetLinkedInAggregateAnalyticsWithResponse()` | Get LinkedIn aggregate stats |
| `client.GetLinkedInPostAnalyticsWithResponse()` | Get LinkedIn post stats |
| `client.GetLinkedInPostReactionsWithResponse()` | Get LinkedIn post reactions |
| `client.GetPostTimelineWithResponse()` | Get post analytics timeline |
| `client.GetPostingFrequencyWithResponse()` | Get posting frequency vs engagement |
| `client.GetYouTubeDailyViewsWithResponse()` | Get YouTube daily views |

### Account Groups
| Method | Description |
|--------|-------------|
| `client.ListAccountGroupsWithResponse()` | List groups |
| `client.CreateAccountGroupWithResponse()` | Create group |
| `client.UpdateAccountGroupWithResponse()` | Update group |
| `client.DeleteAccountGroupWithResponse()` | Delete group |

### Queue
| Method | Description |
|--------|-------------|
| `client.ListQueueSlotsWithResponse()` | List schedules |
| `client.CreateQueueSlotWithResponse()` | Create schedule |
| `client.GetNextQueueSlotWithResponse()` | Get next available slot |
| `client.UpdateQueueSlotWithResponse()` | Update schedule |
| `client.DeleteQueueSlotWithResponse()` | Delete schedule |
| `client.PreviewQueueWithResponse()` | Preview upcoming slots |

### Webhooks
| Method | Description |
|--------|-------------|
| `client.CreateWebhookSettingsWithResponse()` | Create webhook |
| `client.GetWebhookLogsWithResponse()` | Get delivery logs |
| `client.GetWebhookSettingsWithResponse()` | List webhooks |
| `client.UpdateWebhookSettingsWithResponse()` | Update webhook |
| `client.DeleteWebhookSettingsWithResponse()` | Delete webhook |
| `client.TestWebhookWithResponse()` | Send test webhook |

### API Keys
| Method | Description |
|--------|-------------|
| `client.ListApiKeysWithResponse()` | List keys |
| `client.CreateApiKeyWithResponse()` | Create key |
| `client.DeleteApiKeyWithResponse()` | Delete key |

### Media
| Method | Description |
|--------|-------------|
| `client.GetMediaPresignedUrlWithResponse()` | Get presigned upload URL |

### Tools
| Method | Description |
|--------|-------------|
| `client.GetYouTubeTranscriptWithResponse()` | Get YouTube transcript |
| `client.CheckInstagramHashtagsWithResponse()` | Check IG hashtag bans |
| `client.DownloadBlueskyMediaWithResponse()` | Download Bluesky media |
| `client.DownloadFacebookVideoWithResponse()` | Download Facebook video |
| `client.DownloadInstagramMediaWithResponse()` | Download Instagram media |
| `client.DownloadLinkedInVideoWithResponse()` | Download LinkedIn video |
| `client.DownloadTikTokVideoWithResponse()` | Download TikTok video |
| `client.DownloadTwitterMediaWithResponse()` | Download Twitter/X media |
| `client.DownloadYouTubeVideoWithResponse()` | Download YouTube video |

### Users
| Method | Description |
|--------|-------------|
| `client.ListUsersWithResponse()` | List users |
| `client.GetUserWithResponse()` | Get user |

### Usage
| Method | Description |
|--------|-------------|
| `client.GetUsageStatsWithResponse()` | Get plan and usage stats |

### Logs
| Method | Description |
|--------|-------------|
| `client.ListConnectionLogsWithResponse()` | List connection logs |
| `client.ListPostsLogsWithResponse()` | List publishing logs |
| `client.GetPostLogsWithResponse()` | Get post logs |

### Connect (OAuth)
| Method | Description |
|--------|-------------|
| `client.ListFacebookPagesWithResponse()` | List Facebook pages |
| `client.ListGoogleBusinessLocationsWithResponse()` | List GBP locations |
| `client.ListLinkedInOrganizationsWithResponse()` | List LinkedIn orgs |
| `client.ListPinterestBoardsForSelectionWithResponse()` | List Pinterest boards |
| `client.ListSnapchatProfilesWithResponse()` | List Snapchat profiles |
| `client.GetConnectUrlWithResponse()` | Get OAuth connect URL |
| `client.GetFacebookPagesWithResponse()` | List Facebook pages |
| `client.GetGmbLocationsWithResponse()` | List GBP locations |
| `client.GetLinkedInOrganizationsWithResponse()` | List LinkedIn orgs |
| `client.GetPendingOAuthDataWithResponse()` | Get pending OAuth data |
| `client.GetPinterestBoardsWithResponse()` | List Pinterest boards |
| `client.GetRedditFlairsWithResponse()` | List subreddit flairs |
| `client.GetRedditSubredditsWithResponse()` | List Reddit subreddits |
| `client.GetTelegramConnectStatusWithResponse()` | Generate Telegram code |
| `client.UpdateFacebookPageWithResponse()` | Update Facebook page |
| `client.UpdateGmbLocationWithResponse()` | Update GBP location |
| `client.UpdateLinkedInOrganizationWithResponse()` | Switch LinkedIn account type |
| `client.UpdatePinterestBoardsWithResponse()` | Set default Pinterest board |
| `client.UpdateRedditSubredditsWithResponse()` | Set default subreddit |
| `client.CompleteTelegramConnectWithResponse()` | Check Telegram status |
| `client.ConnectBlueskyCredentialsWithResponse()` | Connect Bluesky account |
| `client.ConnectWhatsAppCredentialsWithResponse()` | Connect WhatsApp via credentials |
| `client.HandleOAuthCallbackWithResponse()` | Complete OAuth callback |
| `client.InitiateTelegramConnectWithResponse()` | Connect Telegram directly |
| `client.SelectFacebookPageWithResponse()` | Select Facebook page |
| `client.SelectGoogleBusinessLocationWithResponse()` | Select GBP location |
| `client.SelectLinkedInOrganizationWithResponse()` | Select LinkedIn org |
| `client.SelectPinterestBoardWithResponse()` | Select Pinterest board |
| `client.SelectSnapchatProfileWithResponse()` | Select Snapchat profile |

### Reddit
| Method | Description |
|--------|-------------|
| `client.GetRedditFeedWithResponse()` | Get subreddit feed |
| `client.SearchRedditWithResponse()` | Search posts |

### Account Settings
| Method | Description |
|--------|-------------|
| `client.GetInstagramIceBreakersWithResponse()` | Get IG ice breakers |
| `client.GetMessengerMenuWithResponse()` | Get FB persistent menu |
| `client.GetTelegramCommandsWithResponse()` | Get TG bot commands |
| `client.DeleteInstagramIceBreakersWithResponse()` | Delete IG ice breakers |
| `client.DeleteMessengerMenuWithResponse()` | Delete FB persistent menu |
| `client.DeleteTelegramCommandsWithResponse()` | Delete TG bot commands |
| `client.SetInstagramIceBreakersWithResponse()` | Set IG ice breakers |
| `client.SetMessengerMenuWithResponse()` | Set FB persistent menu |
| `client.SetTelegramCommandsWithResponse()` | Set TG bot commands |

### Comments (Inbox)
| Method | Description |
|--------|-------------|
| `client.ListInboxCommentsWithResponse()` | List commented posts |
| `client.GetInboxPostCommentsWithResponse()` | Get post comments |
| `client.DeleteInboxCommentWithResponse()` | Delete comment |
| `client.HideInboxCommentWithResponse()` | Hide comment |
| `client.LikeInboxCommentWithResponse()` | Like comment |
| `client.ReplyToInboxPostWithResponse()` | Reply to comment |
| `client.SendPrivateReplyToCommentWithResponse()` | Send private reply |
| `client.UnhideInboxCommentWithResponse()` | Unhide comment |
| `client.UnlikeInboxCommentWithResponse()` | Unlike comment |

### GMB Attributes
| Method | Description |
|--------|-------------|
| `client.GetGoogleBusinessAttributesWithResponse()` | Get attributes |
| `client.UpdateGoogleBusinessAttributesWithResponse()` | Update attributes |

### GMB Food Menus
| Method | Description |
|--------|-------------|
| `client.GetGoogleBusinessFoodMenusWithResponse()` | Get food menus |
| `client.UpdateGoogleBusinessFoodMenusWithResponse()` | Update food menus |

### GMB Location Details
| Method | Description |
|--------|-------------|
| `client.GetGoogleBusinessLocationDetailsWithResponse()` | Get location details |
| `client.UpdateGoogleBusinessLocationDetailsWithResponse()` | Update location details |

### GMB Media
| Method | Description |
|--------|-------------|
| `client.ListGoogleBusinessMediaWithResponse()` | List media |
| `client.CreateGoogleBusinessMediaWithResponse()` | Upload photo |
| `client.DeleteGoogleBusinessMediaWithResponse()` | Delete photo |

### GMB Place Actions
| Method | Description |
|--------|-------------|
| `client.ListGoogleBusinessPlaceActionsWithResponse()` | List action links |
| `client.CreateGoogleBusinessPlaceActionWithResponse()` | Create action link |
| `client.DeleteGoogleBusinessPlaceActionWithResponse()` | Delete action link |

### Messages (Inbox)
| Method | Description |
|--------|-------------|
| `client.ListInboxConversationsWithResponse()` | List conversations |
| `client.GetInboxConversationWithResponse()` | Get conversation |
| `client.GetInboxConversationMessagesWithResponse()` | List messages |
| `client.UpdateInboxConversationWithResponse()` | Update conversation status |
| `client.EditInboxMessageWithResponse()` | Edit message |
| `client.SendInboxMessageWithResponse()` | Send message |

### Reviews (Inbox)
| Method | Description |
|--------|-------------|
| `client.ListInboxReviewsWithResponse()` | List reviews |
| `client.DeleteInboxReviewReplyWithResponse()` | Delete review reply |
| `client.ReplyToInboxReviewWithResponse()` | Reply to review |

### Twitter Engagement
| Method | Description |
|--------|-------------|
| `client.BookmarkPostWithResponse()` | Bookmark a tweet |
| `client.FollowUserWithResponse()` | Follow a user |
| `client.RemoveBookmarkWithResponse()` | Remove bookmark |
| `client.RetweetPostWithResponse()` | Retweet a post |
| `client.UndoRetweetWithResponse()` | Undo retweet |
| `client.UnfollowUserWithResponse()` | Unfollow a user |

### Validate
| Method | Description |
|--------|-------------|
| `client.ValidateMediaWithResponse()` | Validate media URL |
| `client.ValidatePostWithResponse()` | Validate post content |
| `client.ValidatePostLengthWithResponse()` | Validate post character count |
| `client.ValidateSubredditWithResponse()` | Check subreddit existence |

### WhatsApp
| Method | Description |
|--------|-------------|
| `client.BulkDeleteWhatsAppContactsWithResponse()` | Bulk delete contacts |
| `client.BulkUpdateWhatsAppContactsWithResponse()` | Bulk update contacts |
| `client.CreateWhatsAppBroadcastWithResponse()` | Create broadcast |
| `client.CreateWhatsAppContactWithResponse()` | Create contact |
| `client.CreateWhatsAppTemplateWithResponse()` | Create template |
| `client.GetWhatsAppBroadcastWithResponse()` | Get broadcast |
| `client.GetWhatsAppBroadcastRecipientsWithResponse()` | List recipients |
| `client.GetWhatsAppBroadcastsWithResponse()` | List broadcasts |
| `client.GetWhatsAppBusinessProfileWithResponse()` | Get business profile |
| `client.GetWhatsAppContactWithResponse()` | Get contact |
| `client.GetWhatsAppContactsWithResponse()` | List contacts |
| `client.GetWhatsAppDisplayNameWithResponse()` | Get display name and review status |
| `client.GetWhatsAppGroupsWithResponse()` | List contact groups |
| `client.GetWhatsAppTemplateWithResponse()` | Get template |
| `client.GetWhatsAppTemplatesWithResponse()` | List templates |
| `client.UpdateWhatsAppBusinessProfileWithResponse()` | Update business profile |
| `client.UpdateWhatsAppContactWithResponse()` | Update contact |
| `client.UpdateWhatsAppDisplayNameWithResponse()` | Request display name change |
| `client.UpdateWhatsAppTemplateWithResponse()` | Update template |
| `client.DeleteWhatsAppBroadcastWithResponse()` | Delete broadcast |
| `client.DeleteWhatsAppContactWithResponse()` | Delete contact |
| `client.DeleteWhatsAppGroupWithResponse()` | Delete group |
| `client.DeleteWhatsAppTemplateWithResponse()` | Delete template |
| `client.AddWhatsAppBroadcastRecipientsWithResponse()` | Add recipients |
| `client.CancelWhatsAppBroadcastScheduleWithResponse()` | Cancel scheduled broadcast |
| `client.ImportWhatsAppContactsWithResponse()` | Bulk import contacts |
| `client.RemoveWhatsAppBroadcastRecipientsWithResponse()` | Remove recipients |
| `client.RenameWhatsAppGroupWithResponse()` | Rename group |
| `client.ScheduleWhatsAppBroadcastWithResponse()` | Schedule broadcast |
| `client.SendWhatsAppBroadcastWithResponse()` | Send broadcast |
| `client.SendWhatsAppBulkWithResponse()` | Bulk send template messages |
| `client.UploadWhatsAppProfilePhotoWithResponse()` | Upload profile picture |

### WhatsApp Phone Numbers
| Method | Description |
|--------|-------------|
| `client.GetWhatsAppPhoneNumberWithResponse()` | Get phone number |
| `client.GetWhatsAppPhoneNumbersWithResponse()` | List phone numbers |
| `client.PurchaseWhatsAppPhoneNumberWithResponse()` | Purchase phone number |
| `client.ReleaseWhatsAppPhoneNumberWithResponse()` | Release phone number |

### Invites
| Method | Description |
|--------|-------------|
| `client.CreateInviteTokenWithResponse()` | Create invite token |

## Documentation

- [API Reference](https://docs.zernio.com/api-reference)
- [Getting Started Guide](https://docs.zernio.com/quickstart)

## License

MIT
