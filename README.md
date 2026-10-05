# Zernio Go SDK

Official Go client library for the [Zernio API](https://docs.zernio.com) - Schedule and manage social media posts across multiple platforms.

## Installation

```bash
go get github.com/zernio-dev/zernio-go
```

> **Upgrading from `v0.0.x`?** `v0.1.0` switches to a new code generator and is
> a breaking change. The import path is unchanged, but client construction and
> method calls differ — see [MIGRATION.md](./MIGRATION.md).

> **Using enum constants?** They are now prefixed with their type
> (`zernio.ADSTATUS_REJECTED`, not `zernio.REJECTED`), see
> [MIGRATION.md](./MIGRATION.md#prefixed-enum-constants).

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
    // Defaults to the https://zernio.com/api base URL from the spec.
    client := zernio.NewAPIClient(zernio.NewConfiguration())

    // Authenticate with your Bearer API key via the request context.
    ctx := context.WithValue(context.Background(), zernio.ContextAccessToken, "YOUR_API_KEY")

    // List accounts
    accounts, _, err := client.AccountsAPI.ListAccounts(ctx).Execute()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Accounts: %+v\n", accounts)
}
```

## SDK Reference

### Posts
| Method | Description |
|--------|-------------|
| `client.PostsAPI.ListPosts(ctx)` | List posts |
| `client.PostsAPI.BulkUploadPosts(ctx)` | Bulk upload from CSV |
| `client.PostsAPI.CreatePost(ctx)` | Create post |
| `client.PostsAPI.GetPost(ctx)` | Get post |
| `client.PostsAPI.UpdatePost(ctx)` | Update post |
| `client.PostsAPI.UpdatePostMetadata(ctx)` | Update post metadata |
| `client.PostsAPI.DeletePost(ctx)` | Delete post |
| `client.PostsAPI.EditPost(ctx)` | Edit published post |
| `client.PostsAPI.RetryPost(ctx)` | Retry failed post |
| `client.PostsAPI.UnpublishPost(ctx)` | Unpublish post |

### Accounts
| Method | Description |
|--------|-------------|
| `client.AccountsAPI.GetAllAccountsHealth(ctx)` | Check accounts health |
| `client.AccountsAPI.ListAccounts(ctx)` | List accounts |
| `client.AccountsAPI.ListBusinessPartners(ctx)` | List partner businesses of the Page |
| `client.AccountsAPI.ListTikTokCommercialMusic(ctx)` | List trending commercial music |
| `client.AccountsAPI.GetAccountHealth(ctx)` | Check account health |
| `client.AccountsAPI.GetAccountPosts(ctx)` | List posts published on the platform |
| `client.AccountsAPI.GetBlueskySettings(ctx)` | Get Bluesky account settings |
| `client.AccountsAPI.GetFollowerStats(ctx)` | Get follower stats |
| `client.GMBReviewsAPI.GetGoogleBusinessReview(ctx)` | Get a review |
| `client.GMBReviewsAPI.GetGoogleBusinessReviews(ctx)` | Get reviews |
| `client.AccountsAPI.GetInstagramFollowStatus(ctx)` | Check whether an Instagram user follows the account |
| `client.LinkedInMentionsAPI.GetLinkedInMentions(ctx)` | Resolve LinkedIn mention |
| `client.AccountsAPI.GetSlackSettings(ctx)` | Get Slack account settings |
| `client.AccountsAPI.GetTikTokCreatorInfo(ctx)` | Get TikTok creator info |
| `client.AccountsAPI.UpdateAccount(ctx)` | Update account |
| `client.AccountsAPI.UpdateBlueskySettings(ctx)` | Update Bluesky account settings |
| `client.AccountsAPI.UpdateSlackSettings(ctx)` | Update Slack account settings |
| `client.AccountsAPI.DeleteAccount(ctx)` | Disconnect account |
| `client.GMBReviewsAPI.DeleteGoogleBusinessReviewReply(ctx)` | Delete a review reply |
| `client.GMBReviewsAPI.BatchGetGoogleBusinessReviews(ctx)` | Batch get reviews |
| `client.AccountsAPI.GrantBusinessPartner(ctx)` | Share the Page with a partner business |
| `client.AccountsAPI.MoveAccountToProfile(ctx)` | Move account to another profile |
| `client.GMBReviewsAPI.ReplyToGoogleBusinessReview(ctx)` | Reply to a review |
| `client.AccountsAPI.RevokeBusinessPartner(ctx)` | Revoke a partner business from the Page |
| `client.AccountsAPI.SearchTikTokLocations(ctx)` | Search TikTok location tags |

### Profiles
| Method | Description |
|--------|-------------|
| `client.ProfilesAPI.ListProfiles(ctx)` | List profiles |
| `client.ProfilesAPI.CreateProfile(ctx)` | Create profile |
| `client.ProfilesAPI.GetProfile(ctx)` | Get profile |
| `client.ProfilesAPI.UpdateProfile(ctx)` | Update profile |
| `client.ProfilesAPI.DeleteProfile(ctx)` | Delete profile |

### Analytics
| Method | Description |
|--------|-------------|
| `client.AnalyticsAPI.GetAnalytics(ctx)` | Get post analytics |
| `client.AnalyticsAPI.GetAnalyticsDashboard(ctx)` | Get an analytics dashboard |
| `client.AnalyticsAPI.GetAnalyticsDelta(ctx)` | Analytics changed since a cursor |
| `client.AnalyticsAPI.GetBestTimeToPost(ctx)` | Get best times to post |
| `client.AnalyticsAPI.GetContentDecay(ctx)` | Get content performance decay |
| `client.AnalyticsAPI.GetDailyMetrics(ctx)` | Get daily aggregated metrics |
| `client.AnalyticsAPI.GetFacebookDemographics(ctx)` | Get Facebook Page demographics |
| `client.AnalyticsAPI.GetFacebookPageInsights(ctx)` | Get Facebook Page insights |
| `client.AnalyticsAPI.GetFacebookPostEarnings(ctx)` | Get Facebook post monetization earnings |
| `client.AnalyticsAPI.GetFacebookPostReactions(ctx)` | Get Facebook post reactions |
| `client.AnalyticsAPI.GetGoogleBusinessPerformance(ctx)` | Get Google Business Profile performance metrics |
| `client.AnalyticsAPI.GetGoogleBusinessSearchKeywords(ctx)` | Get Google Business Profile search keywords |
| `client.AnalyticsAPI.GetInstagramAccountInsights(ctx)` | Get Instagram insights |
| `client.AnalyticsAPI.GetInstagramDemographics(ctx)` | Get Instagram demographics |
| `client.AnalyticsAPI.GetInstagramFollowerHistory(ctx)` | Get Instagram follower history |
| `client.AnalyticsAPI.GetLinkedInAggregateAnalytics(ctx)` | Get LinkedIn aggregate stats |
| `client.AnalyticsAPI.GetLinkedInOrgAggregateAnalytics(ctx)` | Get LinkedIn org analytics |
| `client.AnalyticsAPI.GetLinkedInPostAnalytics(ctx)` | Get LinkedIn post stats |
| `client.AnalyticsAPI.GetLinkedInPostReactions(ctx)` | Get LinkedIn post reactions |
| `client.AnalyticsAPI.GetPostTimeline(ctx)` | Get post analytics timeline |
| `client.AnalyticsAPI.GetPostingFrequency(ctx)` | Get frequency vs engagement |
| `client.AnalyticsAPI.GetTikTokAccountInsights(ctx)` | Get TikTok account-level insights |
| `client.AnalyticsAPI.GetYouTubeChannelInsights(ctx)` | Get YouTube channel insights |
| `client.AnalyticsAPI.GetYouTubeDailyViews(ctx)` | Get YouTube daily views |
| `client.AnalyticsAPI.GetYouTubeDemographics(ctx)` | Get YouTube demographics |
| `client.AnalyticsAPI.GetYouTubeVideoRetention(ctx)` | Get YouTube video retention curve |
| `client.AnalyticsAPI.SyncExternalPosts(ctx)` | Sync an external post |

### Account Groups
| Method | Description |
|--------|-------------|
| `client.AccountGroupsAPI.ListAccountGroups(ctx)` | List groups |
| `client.AccountGroupsAPI.CreateAccountGroup(ctx)` | Create group |
| `client.AccountGroupsAPI.UpdateAccountGroup(ctx)` | Update group |
| `client.AccountGroupsAPI.DeleteAccountGroup(ctx)` | Delete group |

### Queue
| Method | Description |
|--------|-------------|
| `client.QueueAPI.ListQueueSlots(ctx)` | List schedules |
| `client.QueueAPI.CreateQueueSlot(ctx)` | Create schedule |
| `client.QueueAPI.GetNextQueueSlot(ctx)` | Get next available slot |
| `client.QueueAPI.UpdateQueueSlot(ctx)` | Update schedule |
| `client.QueueAPI.DeleteQueueSlot(ctx)` | Delete schedule |
| `client.QueueAPI.PreviewQueue(ctx)` | Preview upcoming slots |

### Webhooks
| Method | Description |
|--------|-------------|
| `client.WebhooksAPI.CreateWebhookSettings(ctx)` | Create webhook |
| `client.WebhooksAPI.GetWebhookLogs(ctx)` | List webhook delivery logs |
| `client.WebhooksAPI.GetWebhookSettings(ctx)` | List webhooks |
| `client.WebhooksAPI.UpdateWebhookSettings(ctx)` | Update webhook |
| `client.WebhooksAPI.DeleteWebhookSettings(ctx)` | Delete webhook |
| `client.WebhooksAPI.RedeliverWebhookEvent(ctx)` | Redeliver a webhook event |
| `client.WebhooksAPI.TestWebhook(ctx)` | Send test webhook |

### API Keys
| Method | Description |
|--------|-------------|
| `client.APIKeysAPI.ListApiKeys(ctx)` | List keys |
| `client.APIKeysAPI.CreateApiKey(ctx)` | Create key |
| `client.APIKeysAPI.DeleteApiKey(ctx)` | Delete key |
| `client.APIKeysAPI.VerifyCredential(ctx)` | Verify credential |

### Media
| Method | Description |
|--------|-------------|
| `client.MediaAPI.GetMediaPresignedUrl(ctx)` | Get upload URL |

### Tools
| Method | Description |
|--------|-------------|
| `client.ToolsAPI.DownloadTikTokVideo(ctx)` | Download a TikTok video |

### Users
| Method | Description |
|--------|-------------|
| `client.UsersAPI.ListUsers(ctx)` | List users |
| `client.UsersAPI.GetUser(ctx)` | Get user |

### Usage
| Method | Description |
|--------|-------------|
| `client.UsageAPI.GetBilling(ctx)` | Account billing snapshot (plan, cycle, balance, caps, status) |
| `client.UsageAPI.GetCallsUsage(ctx)` | Calling usage and cost |
| `client.UsageAPI.GetSmsUsage(ctx)` | SMS usage (volumes) |
| `client.UsageAPI.GetUsage(ctx)` | Usage snapshot (default) or billed-spend metering (with params) |
| `client.UsageAPI.GetUsageStats(ctx)` | Get plan and usage snapshot (plan, limits, payment status) |
| `client.UsageAPI.GetXApiPricing(ctx)` | Get X API pricing table |

### Logs
| Method | Description |
|--------|-------------|
| `client.LogsAPI.ListLogs(ctx)` | List activity logs |

### Connect (OAuth)
| Method | Description |
|--------|-------------|
| `client.ConnectAPI.ListFacebookPages(ctx)` | List Facebook pages |
| `client.ConnectAPI.ListGoogleBusinessLocations(ctx)` | List Google Business Profile locations |
| `client.ConnectAPI.ListInstagramPages(ctx)` | List Pages with a linked Instagram account |
| `client.ConnectAPI.ListLinkedInOrganizations(ctx)` | List LinkedIn orgs |
| `client.ConnectAPI.ListPinterestBoardsForSelection(ctx)` | List Pinterest boards |
| `client.ConnectAPI.ListSlackChannels(ctx)` | List Slack channels for the channel picker |
| `client.ConnectAPI.ListSnapchatProfiles(ctx)` | List Snapchat profiles |
| `client.ConnectAPI.ListWhatsAppPhoneNumbers(ctx)` | List numbers for selection |
| `client.ConnectAPI.CreatePinterestBoard(ctx)` | Create Pinterest board |
| `client.ConnectAPI.CreateYoutubePlaylist(ctx)` | Create YouTube playlist |
| `client.ConnectAPI.GetConnectUrl(ctx)` | Get OAuth connect URL |
| `client.ConnectAPI.GetFacebookPages(ctx)` | List Facebook pages |
| `client.ConnectAPI.GetGmbLocations(ctx)` | List Google Business Profile locations |
| `client.ConnectAPI.GetLinkedInOrganizations(ctx)` | List LinkedIn orgs |
| `client.ConnectAPI.GetPageWebhookSubscription(ctx)` | Read a Facebook Page's webhook subscription |
| `client.ConnectAPI.GetPendingOAuthData(ctx)` | Get pending OAuth data |
| `client.ConnectAPI.GetPinterestBoards(ctx)` | List Pinterest boards |
| `client.ConnectAPI.GetRedditFlairs(ctx)` | List subreddit flairs |
| `client.ConnectAPI.GetRedditSubreddits(ctx)` | List Reddit subreddits |
| `client.ConnectAPI.GetShopifyConnectUrl(ctx)` | Get Shopify OAuth connect URL |
| `client.ConnectAPI.GetSubredditRules(ctx)` | Get subreddit rules |
| `client.ConnectAPI.GetTelegramConnectStatus(ctx)` | Generate Telegram code |
| `client.ConnectAPI.GetWhatsAppSdkConfig(ctx)` | Get Embedded Signup SDK config |
| `client.ConnectAPI.GetWordPressAuthUrl(ctx)` | Get WordPress.com OAuth connect URL |
| `client.ConnectAPI.GetYoutubeCaptions(ctx)` | Get a YouTube video transcript |
| `client.ConnectAPI.GetYoutubePlaylists(ctx)` | List YouTube playlists |
| `client.ConnectAPI.UpdateFacebookPage(ctx)` | Update Facebook page |
| `client.ConnectAPI.UpdateGmbLocation(ctx)` | Update Google Business Profile location |
| `client.ConnectAPI.UpdateLinkedInOrganization(ctx)` | Switch LinkedIn account type |
| `client.ConnectAPI.UpdatePinterestBoards(ctx)` | Set default Pinterest board |
| `client.ConnectAPI.UpdateRedditSubreddits(ctx)` | Set default subreddit |
| `client.ConnectAPI.UpdateYoutubeDefaultPlaylist(ctx)` | Set default YouTube playlist |
| `client.ConnectAPI.AssignGoogleBusinessLocation(ctx)` | Assign Google Business Profile location to another profile |
| `client.ConnectAPI.CompleteMetaAdsBusinessLogin(ctx)` | Complete Meta business login |
| `client.ConnectAPI.CompleteTelegramConnect(ctx)` | Check Telegram status |
| `client.ConnectAPI.CompleteWhatsAppPhoneSelection(ctx)` | Complete number selection |
| `client.ConnectAPI.ConfigureTikTokAdsBrandIdentity(ctx)` | Set TikTok brand identity |
| `client.ConnectAPI.ConnectAds(ctx)` | Connect ads for a platform |
| `client.ConnectAPI.ConnectBlueskyCredentials(ctx)` | Connect Bluesky account |
| `client.ConnectAPI.ConnectDiscordChannel(ctx)` | Connect a Discord channel |
| `client.ConnectAPI.ConnectOpenAIAdsCredentials(ctx)` | Connect an OpenAI Ads account |
| `client.ConnectAPI.ConnectShopifyWithToken(ctx)` | Connect a Shopify store with a custom-app Admin token |
| `client.ConnectAPI.ConnectSlackChannel(ctx)` | Connect a Slack channel |
| `client.ConnectAPI.ConnectWhatsAppCredentials(ctx)` | Connect WhatsApp via credentials |
| `client.ConnectAPI.ConnectWhatsAppEmbeddedSignup(ctx)` | Connect WhatsApp from Embedded Signup |
| `client.ConnectAPI.ConnectWordPressWithApplicationPassword(ctx)` | Connect self-hosted WordPress with an application password |
| `client.ConnectAPI.HandleOAuthCallback(ctx)` | Complete OAuth callback |
| `client.ConnectAPI.InitiateTelegramConnect(ctx)` | Connect Telegram directly |
| `client.ConnectAPI.ResyncPageWebhookSubscription(ctx)` | Re-subscribe a Facebook Page to Zernio's webhooks |
| `client.ConnectAPI.SelectFacebookPage(ctx)` | Select Facebook page |
| `client.ConnectAPI.SelectGoogleBusinessLocation(ctx)` | Select Google Business Profile location |
| `client.ConnectAPI.SelectInstagramAccount(ctx)` | Select the Page whose Instagram account to connect |
| `client.ConnectAPI.SelectLinkedInOrganization(ctx)` | Select LinkedIn org |
| `client.ConnectAPI.SelectPinterestBoard(ctx)` | Select Pinterest board |
| `client.ConnectAPI.SelectSnapchatProfile(ctx)` | Select Snapchat profile |
| `client.ConnectAPI.SetRedditPostFlair(ctx)` | Set Reddit post flair |
| `client.ConnectAPI.VoteRedditThing(ctx)` | Vote on a Reddit post or comment |

### Reddit
| Method | Description |
|--------|-------------|
| `client.RedditSearchAPI.GetRedditFeed(ctx)` | Get subreddit feed |
| `client.RedditSearchAPI.SearchReddit(ctx)` | Search posts |

### Account Settings
| Method | Description |
|--------|-------------|
| `client.AccountSettingsAPI.GetInstagramIceBreakers(ctx)` | Get IG ice breakers |
| `client.AccountSettingsAPI.GetMessengerGetStarted(ctx)` | Get FB Get Started button |
| `client.AccountSettingsAPI.GetMessengerMenu(ctx)` | Get FB persistent menu |
| `client.AccountSettingsAPI.GetTelegramCommands(ctx)` | Get TG bot commands |
| `client.AccountSettingsAPI.DeleteInstagramIceBreakers(ctx)` | Delete IG ice breakers |
| `client.AccountSettingsAPI.DeleteMessengerGetStarted(ctx)` | Delete FB Get Started button |
| `client.AccountSettingsAPI.DeleteMessengerMenu(ctx)` | Delete FB persistent menu |
| `client.AccountSettingsAPI.DeleteTelegramCommands(ctx)` | Delete TG bot commands |
| `client.AccountSettingsAPI.SetInstagramIceBreakers(ctx)` | Set IG ice breakers |
| `client.AccountSettingsAPI.SetMessengerGetStarted(ctx)` | Set FB Get Started button |
| `client.AccountSettingsAPI.SetMessengerMenu(ctx)` | Set FB persistent menu |
| `client.AccountSettingsAPI.SetTelegramCommands(ctx)` | Set TG bot commands |

### Ad Accounts
| Method | Description |
|--------|-------------|
| `client.AdAccountsAPI.ListAccountCallouts(ctx)` | List account callouts |
| `client.AdAccountsAPI.ListAccountSitelinks(ctx)` | List account sitelinks |
| `client.AdAccountsAPI.ListAccountStructuredSnippets(ctx)` | List account snippets |
| `client.AdAccountsAPI.ListAdAccountUsers(ctx)` | Ad account users |
| `client.AdAccountsAPI.ListAdAccounts(ctx)` | List ad accounts |
| `client.AdAccountsAPI.ListAdLabels(ctx)` | List ad labels |
| `client.AdAccountsAPI.ListAdNegativeKeywordLists(ctx)` | List negative keyword lists |
| `client.AdAccountsAPI.ListAdStudies(ctx)` | A/B tests and lift studies |
| `client.AdAccountsAPI.ListAdsBusinessCenters(ctx)` | List TikTok Business Centers |
| `client.AdAccountsAPI.ListAdsInstagramAccounts(ctx)` | List Instagram ad identities |
| `client.AdAccountsAPI.ListAdsInstagramPosts(ctx)` | List Instagram posts to boost |
| `client.AdAccountsAPI.ListAdvertisableApplications(ctx)` | List advertisable apps |
| `client.AdAccountsAPI.ListCustomConversions(ctx)` | List custom conversions |
| `client.AdAccountsAPI.ListHighDemandPeriods(ctx)` | List high-demand periods |
| `client.AdAccountsAPI.ListMetaBusinessUsers(ctx)` | Business users |
| `client.AdAccountsAPI.ListMetaBusinesses(ctx)` | Businesses list |
| `client.AdAccountsAPI.ListPageUsers(ctx)` | Page users of a business |
| `client.AdAccountsAPI.ListTikTokAdPixels(ctx)` | List TikTok ad pixels |
| `client.AdAccountsAPI.ListValueRuleSets(ctx)` | List value rule sets |
| `client.AdAccountsAPI.CreateAdAccount(ctx)` | Create Meta ad account |
| `client.AdAccountsAPI.CreateAdLabel(ctx)` | Create a Google Ads label |
| `client.AdAccountsAPI.CreateAdNegativeKeywordList(ctx)` | Create a negative keyword list |
| `client.AdAccountsAPI.CreateCustomConversion(ctx)` | Create custom conversion |
| `client.AdAccountsAPI.CreateHighDemandPeriod(ctx)` | Schedule a budget increase |
| `client.AdAccountsAPI.CreateValueRuleSet(ctx)` | Create a value rule set |
| `client.AdAccountsAPI.GetAdAccountFinance(ctx)` | Ad account finances |
| `client.AdAccountsAPI.GetAdAccountHierarchy(ctx)` | Get manager account hierarchy |
| `client.AdAccountsAPI.GetAdComments(ctx)` | List comments on an ad |
| `client.AdAccountsAPI.GetAdNegativeKeywordList(ctx)` | Get a negative keyword list |
| `client.AdAccountsAPI.GetAdsActivityLog(ctx)` | Ad account change / audit log |
| `client.AdAccountsAPI.GetDsaDefaults(ctx)` | Get ad account DSA defaults |
| `client.AdAccountsAPI.GetDsaRecommendations(ctx)` | Get DSA recommendations |
| `client.AdAccountsAPI.GetIosFourteenCampaignLimits(ctx)` | Get iOS 14 campaign limits |
| `client.AdAccountsAPI.GetValueRuleSet(ctx)` | Read a value rule set |
| `client.AdAccountsAPI.UpdateAccountCallouts(ctx)` | Update account callouts |
| `client.AdAccountsAPI.UpdateAccountSitelinks(ctx)` | Update account sitelinks |
| `client.AdAccountsAPI.UpdateAccountStructuredSnippets(ctx)` | Update account snippets |
| `client.AdAccountsAPI.UpdateAdAccount(ctx)` | Update ad account settings |
| `client.AdAccountsAPI.UpdateAdAccountManagerLink(ctx)` | Accept, decline, cancel or end a manager link |
| `client.AdAccountsAPI.UpdateAdLabel(ctx)` | Update a Google Ads label |
| `client.AdAccountsAPI.UpdateAdNegativeKeywordList(ctx)` | Rename a negative keyword list |
| `client.AdAccountsAPI.UpdateValueRuleSet(ctx)` | Replace a value rule set |
| `client.AdAccountsAPI.DeleteAdComment(ctx)` | Delete an ad comment |
| `client.AdAccountsAPI.DeleteAdNegativeKeywordList(ctx)` | Delete a negative keyword list |
| `client.AdAccountsAPI.DeleteValueRuleSet(ctx)` | Delete a value rule set |
| `client.AdAccountsAPI.AddAccountCallouts(ctx)` | Add account callouts |
| `client.AdAccountsAPI.AddAccountSitelinks(ctx)` | Add account sitelinks |
| `client.AdAccountsAPI.AddAccountStructuredSnippets(ctx)` | Add account snippets |
| `client.AdAccountsAPI.AssignAdAccountUser(ctx)` | Assign a user to an ad account |
| `client.AdAccountsAPI.AssignPageUser(ctx)` | Assign a user to a Page |
| `client.AdAccountsAPI.AttachAdLabel(ctx)` | Attach a Google Ads label |
| `client.AdAccountsAPI.DetachAdLabel(ctx)` | Detach a Google Ads label |
| `client.AdAccountsAPI.HideAdComment(ctx)` | Hide or unhide an ad comment |
| `client.AdAccountsAPI.InviteAdAccountToManager(ctx)` | Invite a client account to a manager |
| `client.AdAccountsAPI.RemoveAccountCallout(ctx)` | Remove account callout |
| `client.AdAccountsAPI.RemoveAccountSitelink(ctx)` | Remove account sitelink |
| `client.AdAccountsAPI.RemoveAccountStructuredSnippet(ctx)` | Remove account snippet |
| `client.AdAccountsAPI.RemoveAdAccountUser(ctx)` | Remove a user from an ad account |
| `client.AdAccountsAPI.RemoveAdLabel(ctx)` | Remove a Google Ads label |
| `client.AdAccountsAPI.RemovePageUser(ctx)` | Remove a user from a Page |
| `client.AdAccountsAPI.ReplaceAdNegativeKeywordListKeywords(ctx)` | Replace negative list keywords |
| `client.AdAccountsAPI.ReplyToAdComment(ctx)` | Reply to an ad comment |

### Ad Audiences
| Method | Description |
|--------|-------------|
| `client.AdAudiencesAPI.ListAdAudiences(ctx)` | List custom audiences |
| `client.AdAudiencesAPI.CreateAdAudience(ctx)` | Create custom audience |
| `client.AdAudiencesAPI.GetAdAudience(ctx)` | Get audience details |
| `client.AdAudiencesAPI.UpdateAdAudience(ctx)` | Update an audience |
| `client.AdAudiencesAPI.DeleteAdAudience(ctx)` | Delete custom audience |
| `client.AdAudiencesAPI.AddUsersToAdAudience(ctx)` | Add users to audience |
| `client.AdAudiencesAPI.ReplaceAdAudienceCompanies(ctx)` | Replace audience companies |

### Ad Campaigns
| Method | Description |
|--------|-------------|
| `client.AdCampaignsAPI.ListAdCampaigns(ctx)` | List campaigns |
| `client.AdCampaignsAPI.ListAdGroupAssets(ctx)` | List ad-group assets |
| `client.AdCampaignsAPI.ListAdKeywords(ctx)` | List Search keywords |
| `client.AdCampaignsAPI.ListAdSets(ctx)` | List ad sets |
| `client.AdCampaignsAPI.ListAds(ctx)` | List ads |
| `client.AdCampaignsAPI.ListBidStrategies(ctx)` | List portfolio bid strategies |
| `client.AdCampaignsAPI.ListCampaignAssets(ctx)` | List campaign assets |
| `client.AdCampaignsAPI.ListCampaignNegativeKeywordLists(ctx)` | List campaign negative lists |
| `client.AdCampaignsAPI.ListCampaignNegativeKeywords(ctx)` | List campaign-level negative keywords |
| `client.AdCampaignsAPI.ListGoogleAssetGroups(ctx)` | List Performance Max asset groups |
| `client.AdCampaignsAPI.ListGoogleRecommendations(ctx)` | List Google Ads recommendations |
| `client.AdCampaignsAPI.BulkUpdateAdCampaignStatus(ctx)` | Pause or resume many campaigns |
| `client.AdCampaignsAPI.CreateAdCampaign(ctx)` | Create a standalone campaign |
| `client.AdCampaignsAPI.CreateAdSet(ctx)` | Create a standalone ad group |
| `client.AdCampaignsAPI.CreateBidStrategy(ctx)` | Create portfolio bid strategy |
| `client.AdCampaignsAPI.CreateGoogleAssetGroup(ctx)` | Create a Performance Max asset group |
| `client.AdCampaignsAPI.CreateStandaloneAd(ctx)` | Create standalone ad |
| `client.AdCampaignsAPI.GetAd(ctx)` | Get ad details |
| `client.AdCampaignsAPI.GetAdCampaignDetails(ctx)` | Get live campaign details |
| `client.AdCampaignsAPI.GetAdReview(ctx)` | Read the platform's review verdict for an ad |
| `client.AdCampaignsAPI.GetAdSetDetails(ctx)` | Get live ad-set details |
| `client.AdCampaignsAPI.GetAdTree(ctx)` | Get campaign tree |
| `client.AdCampaignsAPI.GetAdsTimeline(ctx)` | Get daily account metrics |
| `client.AdCampaignsAPI.GetCampaignAdSchedule(ctx)` | Read a campaign's ad schedule (dayparting) |
| `client.AdCampaignsAPI.GetCampaignBidding(ctx)` | Read a campaign's current bidding |
| `client.AdCampaignsAPI.GetCampaignConversionGoals(ctx)` | Get campaign conversion goals |
| `client.AdCampaignsAPI.GetCampaignTargeting(ctx)` | Read a Google campaign's device, location, and language targeting |
| `client.AdCampaignsAPI.GetGoogleAssetGroup(ctx)` | Get a Performance Max asset group |
| `client.AdCampaignsAPI.UpdateAd(ctx)` | Update ad |
| `client.AdCampaignsAPI.UpdateAdCampaign(ctx)` | Update a campaign |
| `client.AdCampaignsAPI.UpdateAdCampaignStatus(ctx)` | Pause or resume a campaign |
| `client.AdCampaignsAPI.UpdateAdGroupAssets(ctx)` | Update ad-group assets |
| `client.AdCampaignsAPI.UpdateAdKeyword(ctx)` | Pause or enable a Search keyword |
| `client.AdCampaignsAPI.UpdateAdSet(ctx)` | Update an ad set |
| `client.AdCampaignsAPI.UpdateAdSetStatus(ctx)` | Pause or resume a single ad set |
| `client.AdCampaignsAPI.UpdateAdStatus(ctx)` | Pause or resume a single ad |
| `client.AdCampaignsAPI.UpdateBidStrategy(ctx)` | Update portfolio bid strategy |
| `client.AdCampaignsAPI.UpdateCampaignAdSchedule(ctx)` | Replace a campaign's ad schedule (dayparting) |
| `client.AdCampaignsAPI.UpdateCampaignAssets(ctx)` | Update campaign assets |
| `client.AdCampaignsAPI.UpdateCampaignConversionGoals(ctx)` | Update campaign conversion goals |
| `client.AdCampaignsAPI.UpdateCampaignTargeting(ctx)` | Edit a Google campaign's device, location, or language targeting |
| `client.AdCampaignsAPI.UpdateGoogleAssetGroup(ctx)` | Update a Performance Max asset group |
| `client.AdCampaignsAPI.DeleteAd(ctx)` | Cancel an ad |
| `client.AdCampaignsAPI.DeleteAdCampaign(ctx)` | Delete a campaign |
| `client.AdCampaignsAPI.DeleteAdSet(ctx)` | Delete an ad set |
| `client.AdCampaignsAPI.AddAdKeywords(ctx)` | Add Search ad-group keywords |
| `client.AdCampaignsAPI.ApplyGoogleRecommendations(ctx)` | Apply Google Ads recommendations |
| `client.AdCampaignsAPI.AttachAdGroupAssets(ctx)` | Attach ad-group assets |
| `client.AdCampaignsAPI.AttachCampaignAssets(ctx)` | Attach campaign assets |
| `client.AdCampaignsAPI.BoostPost(ctx)` | Boost post as ad |
| `client.AdCampaignsAPI.DismissGoogleRecommendations(ctx)` | Dismiss Google Ads recommendations |
| `client.AdCampaignsAPI.DuplicateAd(ctx)` | Duplicate an ad |
| `client.AdCampaignsAPI.DuplicateAdCampaign(ctx)` | Duplicate a campaign |
| `client.AdCampaignsAPI.DuplicateAdSet(ctx)` | Duplicate an ad set |
| `client.AdCampaignsAPI.EditGoogleAssetGroupAssets(ctx)` | Link or unlink asset group assets |
| `client.AdCampaignsAPI.RemoveAdGroupAssets(ctx)` | Remove ad-group assets |
| `client.AdCampaignsAPI.RemoveAdKeyword(ctx)` | Remove a Search keyword |
| `client.AdCampaignsAPI.RemoveCampaignAssets(ctx)` | Remove campaign assets |
| `client.AdCampaignsAPI.RemoveGoogleAssetGroup(ctx)` | Remove a Performance Max asset group |
| `client.AdCampaignsAPI.ReplaceCampaignNegativeKeywordLists(ctx)` | Replace campaign negative lists |
| `client.AdCampaignsAPI.ReplaceCampaignNegativeKeywords(ctx)` | Replace campaign-level negative keywords |
| `client.AdCampaignsAPI.ReplaceGoogleListingGroupFilters(ctx)` | Replace an asset group's listing-group tree |

### Ad Creatives
| Method | Description |
|--------|-------------|
| `client.AdCreativesAPI.ListAdCreatives(ctx)` | Creative library |
| `client.AdCreativesAPI.ListAdImages(ctx)` | Ad image library |
| `client.AdCreativesAPI.ListAdVideos(ctx)` | Ad video library |
| `client.AdCreativesAPI.ListAdsTikTokIdentities(ctx)` | List TikTok ad identities |
| `client.AdCreativesAPI.ListPartnershipAdContent(ctx)` | List partnership ad content |
| `client.AdCreativesAPI.ListPartnershipAdPermissions(ctx)` | List partnership permissions |
| `client.AdCreativesAPI.CreateAdCreative(ctx)` | Create a standalone creative |
| `client.AdCreativesAPI.GetAdCreative(ctx)` | Creative details |
| `client.AdCreativesAPI.GetAdMedia(ctx)` | Direct video and image URLs for an ad |
| `client.AdCreativesAPI.GetAdPreviews(ctx)` | Render previews of an existing ad |
| `client.AdCreativesAPI.UpdateAdCreative(ctx)` | Rename a creative |
| `client.AdCreativesAPI.DeleteAdCreative(ctx)` | Delete a creative |
| `client.AdCreativesAPI.DeleteAdVideo(ctx)` | Delete an ad video |
| `client.AdCreativesAPI.GenerateAdPreviews(ctx)` | Render pre-create ad previews |
| `client.AdCreativesAPI.SetPartnershipAdPermission(ctx)` | Set partnership permission |
| `client.AdCreativesAPI.UploadAdImage(ctx)` | Upload an ad image from base64 |
| `client.AdCreativesAPI.UploadAdVideo(ctx)` | Upload an ad video |

### Ad Insights
| Method | Description |
|--------|-------------|
| `client.AdInsightsAPI.ListLocalServicesLeadConversations(ctx)` | List lead conversations |
| `client.AdInsightsAPI.ListLocalServicesLeads(ctx)` | Google Local Services Ads leads |
| `client.AdInsightsAPI.CreateAdInsightsReport(ctx)` | Submit async insights report |
| `client.AdInsightsAPI.GetAdAnalytics(ctx)` | Get ad analytics |
| `client.AdInsightsAPI.GetAdInsightsReport(ctx)` | Poll an async insights report run |
| `client.AdInsightsAPI.GetAdsSearchTerms(ctx)` | Google Ads search terms report |
| `client.AdInsightsAPI.GetCampaignAnalytics(ctx)` | Get campaign analytics |
| `client.AdInsightsAPI.GetTikTokSmartPlusMaterialReport(ctx)` | Per-creative performance inside TikTok Smart+ ads |
| `client.AdInsightsAPI.GenerateKeywordHistoricalMetrics(ctx)` | Get historical keyword metrics |
| `client.AdInsightsAPI.GenerateKeywordIdeas(ctx)` | Generate keyword ideas |
| `client.AdInsightsAPI.QueryAdInsights(ctx)` | Flexible live insights query |

### Ad Library
| Method | Description |
|--------|-------------|
| `client.AdLibraryAPI.SearchAdLibrary(ctx)` | Search the public Ad Library |

### Ad Targeting
| Method | Description |
|--------|-------------|
| `client.AdTargetingAPI.GetLinkedInBidPricing(ctx)` | Suggested bid and budget bounds |
| `client.AdTargetingAPI.GetLinkedInSupplyForecast(ctx)` | Forecast ad delivery |
| `client.AdTargetingAPI.EstimateAdReach(ctx)` | Estimate audience reach |
| `client.AdTargetingAPI.SearchAdInterests(ctx)` | Search targeting interests |
| `client.AdTargetingAPI.SearchAdTargeting(ctx)` | Search targeting options |

### Blogs
| Method | Description |
|--------|-------------|
| `client.BlogsAPI.ListBlogArticles(ctx)` | List blog articles |
| `client.BlogsAPI.ListBlogs(ctx)` | List blogs |
| `client.BlogsAPI.CreateBlog(ctx)` | Create a blog |
| `client.BlogsAPI.CreateBlogArticle(ctx)` | Create a blog article |
| `client.BlogsAPI.GetBlog(ctx)` | Get a blog |
| `client.BlogsAPI.GetBlogArticle(ctx)` | Get a blog article |
| `client.BlogsAPI.UpdateBlog(ctx)` | Update a blog |
| `client.BlogsAPI.UpdateBlogArticle(ctx)` | Update a blog article |
| `client.BlogsAPI.DeleteBlog(ctx)` | Delete a blog |
| `client.BlogsAPI.DeleteBlogArticle(ctx)` | Delete a blog article |

### Branded Calling
| Method | Description |
|--------|-------------|
| `client.BrandedCallingAPI.ListBrandedCallingCallReasons(ctx)` | List pre-approved call reasons |
| `client.BrandedCallingAPI.ListBrandedCallingEnterprises(ctx)` | List registered businesses |
| `client.BrandedCallingAPI.ListBrandedCallingIdentities(ctx)` | List caller identities |
| `client.BrandedCallingAPI.ListBrandedCallingIdentityNumbers(ctx)` | List the numbers on a caller identity |
| `client.BrandedCallingAPI.CreateBrandedCallingEnterprise(ctx)` | Register a business for Branded Calling |
| `client.BrandedCallingAPI.CreateBrandedCallingIdentity(ctx)` | Create a caller identity |
| `client.BrandedCallingAPI.GetBrandedCallingEnterprise(ctx)` | Get a registered business |
| `client.BrandedCallingAPI.GetBrandedCallingIdentity(ctx)` | Get a caller identity |
| `client.BrandedCallingAPI.UpdateBrandedCallingIdentity(ctx)` | Edit or resubmit a caller identity |
| `client.BrandedCallingAPI.DeleteBrandedCallingEnterprise(ctx)` | Delete a registered business |
| `client.BrandedCallingAPI.DeleteBrandedCallingIdentity(ctx)` | Delete a caller identity |
| `client.BrandedCallingAPI.AttachBrandedCallingNumbers(ctx)` | Attach numbers to a verified identity |
| `client.BrandedCallingAPI.ConfirmBrandedCallingAuthorizerEmail(ctx)` | Confirm the authorizer's code |
| `client.BrandedCallingAPI.DetachBrandedCallingNumbers(ctx)` | Detach numbers from an identity |
| `client.BrandedCallingAPI.PreflightBrandedCallingIdentity(ctx)` | Dry-run a caller identity before creating it |
| `client.BrandedCallingAPI.ResendBrandedCallingAuthorizerCode(ctx)` | Resend the authorizer's code |
| `client.BrandedCallingAPI.ShareBrandedCallingIdentityForm(ctx)` | Create a caller identity share link |

### Broadcasts
| Method | Description |
|--------|-------------|
| `client.BroadcastsAPI.ListBroadcastRecipients(ctx)` | List broadcast recipients |
| `client.BroadcastsAPI.ListBroadcasts(ctx)` | List broadcasts |
| `client.BroadcastsAPI.CreateBroadcast(ctx)` | Create broadcast draft |
| `client.BroadcastsAPI.GetBroadcast(ctx)` | Get broadcast details |
| `client.BroadcastsAPI.UpdateBroadcast(ctx)` | Update broadcast |
| `client.BroadcastsAPI.DeleteBroadcast(ctx)` | Delete broadcast |
| `client.BroadcastsAPI.AddBroadcastRecipients(ctx)` | Add recipients to a broadcast |
| `client.BroadcastsAPI.CancelBroadcast(ctx)` | Cancel broadcast |
| `client.BroadcastsAPI.ScheduleBroadcast(ctx)` | Schedule broadcast for later |
| `client.BroadcastsAPI.SendBroadcast(ctx)` | Send broadcast now |

### Business Agent
| Method | Description |
|--------|-------------|
| `client.BusinessAgentAPI.ListBusinessAgentAllowlist(ctx)` | List allowlisted consumers |
| `client.BusinessAgentAPI.ListBusinessAgentConnectorTools(ctx)` | List connector tools |
| `client.BusinessAgentAPI.ListBusinessAgentConnectors(ctx)` | List connectors |
| `client.BusinessAgentAPI.ListBusinessAgentFaqs(ctx)` | List FAQs |
| `client.BusinessAgentAPI.ListBusinessAgentFiles(ctx)` | List knowledge files |
| `client.BusinessAgentAPI.ListBusinessAgentSettings(ctx)` | List agent settings |
| `client.BusinessAgentAPI.ListBusinessAgentSkills(ctx)` | List skills |
| `client.BusinessAgentAPI.ListBusinessAgentUiSkills(ctx)` | List UI skills |
| `client.BusinessAgentAPI.ListBusinessAgentWebsites(ctx)` | List crawled websites |
| `client.BusinessAgentAPI.CreateBusinessAgentConnector(ctx)` | Create a connector |
| `client.BusinessAgentAPI.CreateBusinessAgentConnectorTool(ctx)` | Create a connector tool |
| `client.BusinessAgentAPI.CreateBusinessAgentFaq(ctx)` | Create a FAQ |
| `client.BusinessAgentAPI.CreateBusinessAgentSkill(ctx)` | Create a skill |
| `client.BusinessAgentAPI.CreateBusinessAgentUiSkill(ctx)` | Create a UI skill |
| `client.BusinessAgentAPI.GetBusinessAgentBudget(ctx)` | Get usage budgets |
| `client.BusinessAgentAPI.GetBusinessAgentBusinessInformation(ctx)` | Get business information |
| `client.BusinessAgentAPI.GetBusinessAgentConnector(ctx)` | Get a connector |
| `client.BusinessAgentAPI.GetBusinessAgentConnectorLogs(ctx)` | Get connector failure logs |
| `client.BusinessAgentAPI.GetBusinessAgentConnectorTool(ctx)` | Get a connector tool |
| `client.BusinessAgentAPI.GetBusinessAgentEvent(ctx)` | Get a business event status |
| `client.BusinessAgentAPI.GetBusinessAgentFaq(ctx)` | Get a FAQ |
| `client.BusinessAgentAPI.GetBusinessAgentFile(ctx)` | Get a knowledge file |
| `client.BusinessAgentAPI.GetBusinessAgentSkill(ctx)` | Get a skill |
| `client.BusinessAgentAPI.GetBusinessAgentStatus(ctx)` | Get agent setup status |
| `client.BusinessAgentAPI.GetBusinessAgentUiSkill(ctx)` | Get a UI skill |
| `client.BusinessAgentAPI.GetBusinessAgentWebsite(ctx)` | Get a crawled website |
| `client.BusinessAgentAPI.UpdateBusinessAgentConnector(ctx)` | Update a connector |
| `client.BusinessAgentAPI.UpdateBusinessAgentConnectorTool(ctx)` | Update a connector tool |
| `client.BusinessAgentAPI.UpdateBusinessAgentFaq(ctx)` | Update a FAQ |
| `client.BusinessAgentAPI.UpdateBusinessAgentSettings(ctx)` | Update agent settings |
| `client.BusinessAgentAPI.UpdateBusinessAgentSkill(ctx)` | Update a skill |
| `client.BusinessAgentAPI.UpdateBusinessAgentUiSkill(ctx)` | Update a UI skill |
| `client.BusinessAgentAPI.UpdateBusinessAgentWebsite(ctx)` | Update a crawled website |
| `client.BusinessAgentAPI.DeleteBusinessAgentConnector(ctx)` | Delete a connector |
| `client.BusinessAgentAPI.DeleteBusinessAgentConnectorTool(ctx)` | Delete a connector tool |
| `client.BusinessAgentAPI.DeleteBusinessAgentFaq(ctx)` | Delete a FAQ |
| `client.BusinessAgentAPI.DeleteBusinessAgentFile(ctx)` | Delete a knowledge file |
| `client.BusinessAgentAPI.DeleteBusinessAgentSkill(ctx)` | Delete a skill |
| `client.BusinessAgentAPI.DeleteBusinessAgentUiSkill(ctx)` | Delete a UI skill |
| `client.BusinessAgentAPI.DeleteBusinessAgentWebsite(ctx)` | Remove a crawled website |
| `client.BusinessAgentAPI.AddBusinessAgentAllowlistEntry(ctx)` | Allowlist a consumer |
| `client.BusinessAgentAPI.AddBusinessAgentWebsite(ctx)` | Add a website to crawl |
| `client.BusinessAgentAPI.OnboardBusinessAgent(ctx)` | Create the agent |
| `client.BusinessAgentAPI.ReadBusinessAgentEvals(ctx)` | Read evaluation data |
| `client.BusinessAgentAPI.RefreshBusinessAgentConnectorTools(ctx)` | Refresh MCP connector tools |
| `client.BusinessAgentAPI.RemoveBusinessAgentAllowlistEntry(ctx)` | Remove an allowlisted consumer |
| `client.BusinessAgentAPI.ReplaceBusinessAgentBudget(ctx)` | Replace usage budgets |
| `client.BusinessAgentAPI.ReplaceBusinessAgentBusinessInformation(ctx)` | Replace business information |
| `client.BusinessAgentAPI.ResetBusinessAgentBusinessInformation(ctx)` | Reset business information |
| `client.BusinessAgentAPI.RunBusinessAgentConnectorTool(ctx)` | Run a connector tool once |
| `client.BusinessAgentAPI.SendBusinessAgentEvent(ctx)` | Send a business event |
| `client.BusinessAgentAPI.SendBusinessAgentTestMessage(ctx)` | Send a test message |
| `client.BusinessAgentAPI.SetBusinessAgentConnectorCredentials(ctx)` | Set connector credentials |
| `client.BusinessAgentAPI.StartBusinessAgentEvalRun(ctx)` | Start an evaluation run |
| `client.BusinessAgentAPI.UploadBusinessAgentFile(ctx)` | Upload a knowledge file |

### Calls
| Method | Description |
|--------|-------------|
| `client.CallsAPI.ListCalls(ctx)` | List all calls (unified history) |
| `client.CallsAPI.GetCall(ctx)` | Get a call (any channel) |
| `client.CallsAPI.GetCallRecording(ctx)` | Get a call recording |

### Changelog
| Method | Description |
|--------|-------------|
| `client.ChangelogAPI.ListChangelog(ctx)` | List API changelog entries |

### Comment Automations
| Method | Description |
|--------|-------------|
| `client.CommentAutomationsAPI.ListCommentAutomationLogs(ctx)` | List automation logs |
| `client.CommentAutomationsAPI.ListCommentAutomations(ctx)` | List comment-to-DM automations |
| `client.CommentAutomationsAPI.CreateCommentAutomation(ctx)` | Create comment-to-DM automation |
| `client.CommentAutomationsAPI.GetCommentAutomation(ctx)` | Get automation details |
| `client.CommentAutomationsAPI.UpdateCommentAutomation(ctx)` | Update automation settings |
| `client.CommentAutomationsAPI.DeleteCommentAutomation(ctx)` | Delete automation |

### Comments (Inbox)
| Method | Description |
|--------|-------------|
| `client.CommentsAPI.ListInboxComments(ctx)` | List commented posts |
| `client.CommentsAPI.GetInboxPostComments(ctx)` | Get post comments |
| `client.CommentsAPI.DeleteInboxComment(ctx)` | Delete comment |
| `client.CommentsAPI.EditInboxComment(ctx)` | Edit comment |
| `client.CommentsAPI.HideInboxComment(ctx)` | Hide comment |
| `client.CommentsAPI.LikeInboxComment(ctx)` | Like comment |
| `client.CommentsAPI.LikePost(ctx)` | Like post |
| `client.CommentsAPI.PinInboxComment(ctx)` | Pin comment |
| `client.CommentsAPI.ReplyToInboxPost(ctx)` | Reply to comment |
| `client.CommentsAPI.SendPrivateReplyToComment(ctx)` | Send private reply |
| `client.CommentsAPI.SetCommentModeration(ctx)` | Set comment moderation status |
| `client.CommentsAPI.UnhideInboxComment(ctx)` | Unhide comment |
| `client.CommentsAPI.UnlikeInboxComment(ctx)` | Unlike comment |
| `client.CommentsAPI.UnlikePost(ctx)` | Unlike post |
| `client.CommentsAPI.UnpinInboxComment(ctx)` | Unpin comment |

### Commerce
| Method | Description |
|--------|-------------|
| `client.CommerceAPI.ListCommerceCatalogSyncs(ctx)` | List catalog syncs |
| `client.CommerceAPI.ListCommerceChannels(ctx)` | List sales channels |
| `client.CommerceAPI.ListCommerceCollectionMetafields(ctx)` | List collection metafields |
| `client.CommerceAPI.ListCommerceCollections(ctx)` | List collections |
| `client.CommerceAPI.ListCommerceDiscounts(ctx)` | List discounts |
| `client.CommerceAPI.ListCommerceInventory(ctx)` | Get a product's stock |
| `client.CommerceAPI.ListCommerceLocations(ctx)` | List locations |
| `client.CommerceAPI.ListCommerceMarkets(ctx)` | List markets |
| `client.CommerceAPI.ListCommerceMenus(ctx)` | List navigation menus |
| `client.CommerceAPI.ListCommerceMetaobjectDefinitions(ctx)` | List metaobject definitions |
| `client.CommerceAPI.ListCommerceMetaobjects(ctx)` | List metaobjects of a type |
| `client.CommerceAPI.ListCommercePages(ctx)` | List pages |
| `client.CommerceAPI.ListCommercePriceLists(ctx)` | List price lists |
| `client.CommerceAPI.ListCommerceProductMetafields(ctx)` | List product metafields |
| `client.CommerceAPI.ListCommerceProducts(ctx)` | List products |
| `client.CommerceAPI.ListCommerceRedirects(ctx)` | List URL redirects |
| `client.CommerceAPI.CreateCommerceCatalogSync(ctx)` | Sync a store into a Meta catalog |
| `client.CommerceAPI.CreateCommerceCollection(ctx)` | Create a collection |
| `client.CommerceAPI.CreateCommerceDiscount(ctx)` | Create a discount |
| `client.CommerceAPI.CreateCommerceMenu(ctx)` | Create a navigation menu |
| `client.CommerceAPI.CreateCommerceMetaobject(ctx)` | Create a metaobject |
| `client.CommerceAPI.CreateCommercePage(ctx)` | Create a page |
| `client.CommerceAPI.CreateCommerceProduct(ctx)` | Create a product |
| `client.CommerceAPI.CreateCommerceProductOptions(ctx)` | Add options |
| `client.CommerceAPI.CreateCommerceProductVariants(ctx)` | Add variants |
| `client.CommerceAPI.CreateCommerceRedirect(ctx)` | Create a URL redirect |
| `client.CommerceAPI.GetCommerceCatalogSync(ctx)` | Get a catalog sync |
| `client.CommerceAPI.GetCommerceCollection(ctx)` | Get a collection |
| `client.CommerceAPI.GetCommerceDiscount(ctx)` | Get a discount |
| `client.CommerceAPI.GetCommerceMenu(ctx)` | Get a navigation menu |
| `client.CommerceAPI.GetCommerceMetaobject(ctx)` | Get a metaobject |
| `client.CommerceAPI.GetCommercePage(ctx)` | Get a page |
| `client.CommerceAPI.GetCommerceProduct(ctx)` | Get a product |
| `client.CommerceAPI.GetCommerceStore(ctx)` | Get a store |
| `client.CommerceAPI.UpdateCommerceCollection(ctx)` | Update a collection |
| `client.CommerceAPI.UpdateCommerceDiscount(ctx)` | Update a discount |
| `client.CommerceAPI.UpdateCommerceMenu(ctx)` | Replace a navigation menu |
| `client.CommerceAPI.UpdateCommerceMetaobject(ctx)` | Update a metaobject |
| `client.CommerceAPI.UpdateCommercePage(ctx)` | Update a page |
| `client.CommerceAPI.UpdateCommerceProduct(ctx)` | Update a product |
| `client.CommerceAPI.UpdateCommerceProductPrices(ctx)` | Update variant prices |
| `client.CommerceAPI.UpdateCommerceRedirect(ctx)` | Update a URL redirect |
| `client.CommerceAPI.DeleteCommerceCatalogSync(ctx)` | Stop a catalog sync |
| `client.CommerceAPI.DeleteCommerceCollection(ctx)` | Delete a collection |
| `client.CommerceAPI.DeleteCommerceCollectionMetafields(ctx)` | Delete collection metafields |
| `client.CommerceAPI.DeleteCommerceDiscount(ctx)` | Delete a discount |
| `client.CommerceAPI.DeleteCommerceMarketingActivity(ctx)` | Delete a marketing activity |
| `client.CommerceAPI.DeleteCommerceMenu(ctx)` | Delete a navigation menu |
| `client.CommerceAPI.DeleteCommerceMetaobject(ctx)` | Delete a metaobject |
| `client.CommerceAPI.DeleteCommercePage(ctx)` | Delete a page |
| `client.CommerceAPI.DeleteCommercePriceListPrices(ctx)` | Remove fixed prices |
| `client.CommerceAPI.DeleteCommerceProductMetafields(ctx)` | Delete product metafields |
| `client.CommerceAPI.DeleteCommerceProductOptions(ctx)` | Delete options |
| `client.CommerceAPI.DeleteCommerceProductVariants(ctx)` | Delete variants |
| `client.CommerceAPI.DeleteCommerceRedirect(ctx)` | Delete a URL redirect |
| `client.CommerceAPI.AddCommerceDiscountCodes(ctx)` | Add codes to a discount |
| `client.CommerceAPI.AddCommerceMarketingEngagement(ctx)` | Report daily engagement |
| `client.CommerceAPI.AddCommerceProductImages(ctx)` | Add images |
| `client.CommerceAPI.ChangeCommerceCollectionChannels(ctx)` | Publish or unpublish a collection |
| `client.CommerceAPI.ChangeCommerceCollectionProducts(ctx)` | Add or remove products in a collection |
| `client.CommerceAPI.ChangeCommerceInventory(ctx)` | Set or adjust stock |
| `client.CommerceAPI.ChangeCommerceProductChannels(ctx)` | Publish or unpublish a product |
| `client.CommerceAPI.ChangeCommerceProductState(ctx)` | Activate, deactivate, archive or delete products |
| `client.CommerceAPI.ChangeCommerceProductTags(ctx)` | Add or remove tags in bulk |
| `client.CommerceAPI.DuplicateCommerceProduct(ctx)` | Duplicate a product |
| `client.CommerceAPI.RemoveCommerceProductImages(ctx)` | Remove images |
| `client.CommerceAPI.ReorderCommerceCollectionProducts(ctx)` | Reorder products in a collection |
| `client.CommerceAPI.ReorderCommerceProductImages(ctx)` | Reorder images |
| `client.CommerceAPI.RunCommerceCatalogSync(ctx)` | Run a catalog sync now |
| `client.CommerceAPI.SetCommerceCollectionMetafields(ctx)` | Set collection metafields |
| `client.CommerceAPI.SetCommerceDiscountActive(ctx)` | Activate or deactivate a discount |
| `client.CommerceAPI.SetCommercePriceListPrices(ctx)` | Set fixed prices |
| `client.CommerceAPI.SetCommerceProductMetafields(ctx)` | Set product metafields |
| `client.CommerceAPI.UpsertCommerceMarketingActivity(ctx)` | Record a marketing activity |

### Connected Apps
| Method | Description |
|--------|-------------|
| `client.ConnectedAppsAPI.ListConnectedApps(ctx)` | List connected apps |
| `client.ConnectedAppsAPI.RevokeConnectedApp(ctx)` | Revoke connected app |

### Contacts
| Method | Description |
|--------|-------------|
| `client.ContactsAPI.ListContacts(ctx)` | List contacts |
| `client.ContactsAPI.BulkCreateContacts(ctx)` | Bulk create contacts |
| `client.ContactsAPI.CreateContact(ctx)` | Create contact |
| `client.ContactsAPI.GetContact(ctx)` | Get contact |
| `client.ContactsAPI.GetContactChannels(ctx)` | List channels for a contact |
| `client.ContactsAPI.UpdateContact(ctx)` | Update contact |
| `client.ContactsAPI.DeleteContact(ctx)` | Delete contact |

### Conversions
| Method | Description |
|--------|-------------|
| `client.ConversionsAPI.ListAdConversionGoals(ctx)` | List account conversion goals |
| `client.ConversionsAPI.ListConversionActions(ctx)` | List conversion actions |
| `client.ConversionsAPI.ListConversionAssociations(ctx)` | List associated campaigns |
| `client.ConversionsAPI.ListConversionDestinations(ctx)` | List conversion destinations |
| `client.ConversionsAPI.ListCustomConversionGoals(ctx)` | List custom conversion goals |
| `client.ConversionsAPI.CreateConversionAction(ctx)` | Create website conversion action |
| `client.ConversionsAPI.CreateConversionDestination(ctx)` | Create a conversion destination |
| `client.ConversionsAPI.CreateCustomConversionGoal(ctx)` | Create a custom conversion goal |
| `client.ConversionsAPI.GetConversionDestination(ctx)` | Get a conversion destination |
| `client.ConversionsAPI.GetConversionMetrics(ctx)` | Get attribution metrics |
| `client.ConversionsAPI.GetConversionsQuality(ctx)` | Get Event Match Quality |
| `client.ConversionsAPI.UpdateAdConversionGoals(ctx)` | Update account conversion goals |
| `client.ConversionsAPI.UpdateConversionAction(ctx)` | Set a conversion action primary or secondary |
| `client.ConversionsAPI.UpdateConversionDestination(ctx)` | Update a conversion destination |
| `client.ConversionsAPI.UpdateCustomConversionGoal(ctx)` | Update a custom conversion goal |
| `client.ConversionsAPI.DeleteConversionDestination(ctx)` | Delete a conversion destination |
| `client.ConversionsAPI.AddConversionAssociations(ctx)` | Associate campaigns |
| `client.ConversionsAPI.AdjustConversions(ctx)` | Adjust uploaded conversions |
| `client.ConversionsAPI.RemoveConversionAssociations(ctx)` | Remove associated campaigns |
| `client.ConversionsAPI.RemoveCustomConversionGoal(ctx)` | Remove a custom conversion goal |
| `client.ConversionsAPI.SendConversions(ctx)` | Send conversion events |

### Custom Fields
| Method | Description |
|--------|-------------|
| `client.CustomFieldsAPI.ListCustomFields(ctx)` | List custom field definitions |
| `client.CustomFieldsAPI.CreateCustomField(ctx)` | Create custom field |
| `client.CustomFieldsAPI.UpdateCustomField(ctx)` | Update custom field |
| `client.CustomFieldsAPI.DeleteCustomField(ctx)` | Delete custom field |
| `client.CustomFieldsAPI.ClearContactFieldValue(ctx)` | Clear custom field value |
| `client.CustomFieldsAPI.SetContactFieldValue(ctx)` | Set custom field value |

### Discord
| Method | Description |
|--------|-------------|
| `client.DiscordAPI.ListDiscordGuildMembers(ctx)` | List Discord guild members |
| `client.DiscordAPI.ListDiscordGuildRoles(ctx)` | List Discord guild roles |
| `client.DiscordAPI.ListDiscordPinnedMessages(ctx)` | List pinned messages |
| `client.DiscordAPI.ListDiscordScheduledEvents(ctx)` | List Discord scheduled events |
| `client.DiscordAPI.CreateDiscordGuildRole(ctx)` | Create a Discord guild role |
| `client.DiscordAPI.CreateDiscordScheduledEvent(ctx)` | Create a Discord scheduled event |
| `client.DiscordAPI.CreateDiscordThread(ctx)` | Create a Discord public thread |
| `client.DiscordAPI.GetDiscordChannels(ctx)` | List Discord guild channels |
| `client.DiscordAPI.GetDiscordGuildMember(ctx)` | Get a Discord guild member |
| `client.DiscordAPI.GetDiscordScheduledEvent(ctx)` | Get a Discord scheduled event |
| `client.DiscordAPI.GetDiscordSettings(ctx)` | Get Discord account settings |
| `client.DiscordAPI.UpdateDiscordScheduledEvent(ctx)` | Update a Discord scheduled event |
| `client.DiscordAPI.UpdateDiscordSettings(ctx)` | Update Discord settings |
| `client.DiscordAPI.DeleteDiscordGuildRole(ctx)` | Delete a Discord guild role |
| `client.DiscordAPI.DeleteDiscordMessage(ctx)` | Delete a Discord channel message |
| `client.DiscordAPI.DeleteDiscordScheduledEvent(ctx)` | Delete a Discord scheduled event |
| `client.DiscordAPI.AddDiscordMemberRole(ctx)` | Assign a role to a guild member |
| `client.DiscordAPI.CrosspostDiscordMessage(ctx)` | Crosspost Discord message |
| `client.DiscordAPI.EditDiscordGuildRole(ctx)` | Edit a Discord guild role |
| `client.DiscordAPI.PinDiscordMessage(ctx)` | Pin a Discord message |
| `client.DiscordAPI.RemoveDiscordMemberRole(ctx)` | Remove a role from a guild member |
| `client.DiscordAPI.SearchDiscordGuildMembers(ctx)` | Search Discord guild members |
| `client.DiscordAPI.SendDiscordDirectMessage(ctx)` | Send a Discord Direct Message |
| `client.DiscordAPI.UnpinDiscordMessage(ctx)` | Unpin a Discord message |

### Feedback
| Method | Description |
|--------|-------------|
| `client.FeedbackAPI.SubmitFeedback(ctx)` | Submit feedback |

### GMB Attributes
| Method | Description |
|--------|-------------|
| `client.GMBAttributesAPI.GetGmbAttributeMetadata(ctx)` | Get attribute metadata |
| `client.GMBAttributesAPI.GetGoogleBusinessAttributes(ctx)` | Get attributes |
| `client.GMBAttributesAPI.UpdateGoogleBusinessAttributes(ctx)` | Update attributes |

### GMB Food Menus
| Method | Description |
|--------|-------------|
| `client.GMBFoodMenusAPI.GetGoogleBusinessFoodMenus(ctx)` | Get food menus |
| `client.GMBFoodMenusAPI.UpdateGoogleBusinessFoodMenus(ctx)` | Update food menus |

### GMB Location Details
| Method | Description |
|--------|-------------|
| `client.GMBLocationDetailsAPI.GetGoogleBusinessLocationDetails(ctx)` | Get location details |
| `client.GMBLocationDetailsAPI.UpdateGoogleBusinessLocationDetails(ctx)` | Update location details |

### GMB Media
| Method | Description |
|--------|-------------|
| `client.GMBMediaAPI.ListGoogleBusinessMedia(ctx)` | List media |
| `client.GMBMediaAPI.CreateGoogleBusinessMedia(ctx)` | Upload photo |
| `client.GMBMediaAPI.DeleteGoogleBusinessMedia(ctx)` | Delete photo |

### GMB Place Actions
| Method | Description |
|--------|-------------|
| `client.GMBPlaceActionsAPI.ListGoogleBusinessPlaceActions(ctx)` | List action links |
| `client.GMBPlaceActionsAPI.CreateGoogleBusinessPlaceAction(ctx)` | Create action link |
| `client.GMBPlaceActionsAPI.UpdateGoogleBusinessPlaceAction(ctx)` | Update action link |
| `client.GMBPlaceActionsAPI.DeleteGoogleBusinessPlaceAction(ctx)` | Delete action link |

### GMB Services
| Method | Description |
|--------|-------------|
| `client.GMBServicesAPI.GetGoogleBusinessServices(ctx)` | Get services |
| `client.GMBServicesAPI.UpdateGoogleBusinessServices(ctx)` | Replace services |

### GMB Verifications
| Method | Description |
|--------|-------------|
| `client.GMBVerificationsAPI.GetGoogleBusinessVerifications(ctx)` | Get verification state |
| `client.GMBVerificationsAPI.CompleteGoogleBusinessVerification(ctx)` | Complete a verification |
| `client.GMBVerificationsAPI.FetchGoogleBusinessVerificationOptions(ctx)` | Fetch verification options |
| `client.GMBVerificationsAPI.StartGoogleBusinessVerification(ctx)` | Start a verification |

### Inbox Analytics
| Method | Description |
|--------|-------------|
| `client.InboxAnalyticsAPI.ListInboxConversationAnalytics(ctx)` | List conversation analytics |
| `client.InboxAnalyticsAPI.GetInboxConversationAnalytics(ctx)` | Get conversation analytics |
| `client.InboxAnalyticsAPI.GetInboxHeatmap(ctx)` | Get day × hour heatmap |
| `client.InboxAnalyticsAPI.GetInboxResponseTime(ctx)` | Get inbox response-time stats |
| `client.InboxAnalyticsAPI.GetInboxSourceBreakdown(ctx)` | Get inbox source breakdown |
| `client.InboxAnalyticsAPI.GetInboxTopAccounts(ctx)` | Get top accounts by inbox volume |
| `client.InboxAnalyticsAPI.GetInboxVolume(ctx)` | Get inbox messaging volume |

### Instagram
| Method | Description |
|--------|-------------|
| `client.InstagramAPI.ListInstagramStories(ctx)` | List active Instagram stories |
| `client.InstagramAPI.GetInstagramAudio(ctx)` | Get Instagram audio metadata |
| `client.InstagramAPI.GetInstagramBusinessDiscovery(ctx)` | Look up a public Instagram Business account |
| `client.InstagramAPI.GetInstagramPublishingLimit(ctx)` | Get Instagram publishing limit |
| `client.InstagramAPI.GetInstagramStoryInsights(ctx)` | Get Instagram story insights |
| `client.InstagramAPI.SearchInstagramAudio(ctx)` | Search Instagram audio |

### Lead Gen
| Method | Description |
|--------|-------------|
| `client.LeadGenAPI.ListFormLeads(ctx)` | List leads for a single form |
| `client.LeadGenAPI.ListLeadForms(ctx)` | List lead forms |
| `client.LeadGenAPI.ListLeads(ctx)` | List submitted leads |
| `client.LeadGenAPI.CreateLeadForm(ctx)` | Create a lead form |
| `client.LeadGenAPI.CreateTestLead(ctx)` | Create a test lead |
| `client.LeadGenAPI.GetLeadForm(ctx)` | Get a lead form |
| `client.LeadGenAPI.DeleteTestLead(ctx)` | Delete a test lead |
| `client.LeadGenAPI.ArchiveLeadForm(ctx)` | Archive a lead form |

### Mentions
| Method | Description |
|--------|-------------|
| `client.MentionsAPI.ListInboxMentions(ctx)` | List mentions |
| `client.MentionsAPI.ReplyToMention(ctx)` | Reply to a mention |

### Messages (Inbox)
| Method | Description |
|--------|-------------|
| `client.MessagesAPI.ListInboxConversations(ctx)` | List conversations |
| `client.MessagesAPI.CreateInboxConversation(ctx)` | Create conversation |
| `client.MessagesAPI.GetInboxConversation(ctx)` | Get conversation |
| `client.MessagesAPI.GetInboxConversationMessages(ctx)` | List messages |
| `client.MessagesAPI.GetMessageAttachment(ctx)` | Resolve message attachment |
| `client.MessagesAPI.UpdateInboxConversation(ctx)` | Update conversation status |
| `client.MessagesAPI.DeleteInboxMessage(ctx)` | Delete message |
| `client.MessagesAPI.AddMessageReaction(ctx)` | Add reaction |
| `client.MessagesAPI.EditInboxMessage(ctx)` | Edit message |
| `client.MessagesAPI.MarkConversationRead(ctx)` | Mark a conversation as read |
| `client.MessagesAPI.RemoveMessageReaction(ctx)` | Remove reaction |
| `client.MessagesAPI.SearchInboxConversations(ctx)` | Search conversations |
| `client.MessagesAPI.SendInboxMessage(ctx)` | Send message |
| `client.MessagesAPI.SendTypingIndicator(ctx)` | Send typing indicator |
| `client.MessagesAPI.SetConversationThreadControl(ctx)` | Hand a conversation to or from Meta Business Agent |
| `client.MessagesAPI.UploadMediaDirect(ctx)` | Upload media file |

### Messaging Ads
| Method | Description |
|--------|-------------|
| `client.MessagingAdsAPI.CreateCallAd(ctx)` | Create Click-to-Call ad |
| `client.MessagingAdsAPI.CreateCtwaAd(ctx)` | Create CTWA ad (deprecated) |
| `client.MessagingAdsAPI.CreateMessagingAd(ctx)` | Create messaging ad |

### Phone Numbers
| Method | Description |
|--------|-------------|
| `client.PhoneNumbersAPI.ListPhoneNumberCountries(ctx)` | List offerable number countries |
| `client.PhoneNumbersAPI.ListPhoneNumberPortIns(ctx)` | List port-in orders |
| `client.PhoneNumbersAPI.ListPhoneNumberStockWatches(ctx)` | List stock watches |
| `client.PhoneNumbersAPI.ListPhoneNumbers(ctx)` | List phone numbers |
| `client.PhoneNumbersAPI.CreatePhoneNumberKycLink(ctx)` | Create a hosted KYC link |
| `client.PhoneNumbersAPI.CreatePhoneNumberPortIn(ctx)` | Port numbers in |
| `client.PhoneNumbersAPI.CreatePhoneNumberStockWatch(ctx)` | Watch an out-of-stock country |
| `client.PhoneNumbersAPI.GetPhoneNumber(ctx)` | Get phone number |
| `client.PhoneNumbersAPI.GetPhoneNumberClaim(ctx)` | Resolve a number claim |
| `client.PhoneNumbersAPI.GetPhoneNumberKycForm(ctx)` | Get KYC form spec |
| `client.PhoneNumbersAPI.GetPhoneNumberPortClaim(ctx)` | Resolve a port claim |
| `client.PhoneNumbersAPI.GetPhoneNumberPortInOrderRequirements(ctx)` | A port-in order's pending requirements |
| `client.PhoneNumbersAPI.GetPhoneNumberPortInRequirements(ctx)` | Country porting requirements |
| `client.PhoneNumbersAPI.GetPhoneNumberRemediation(ctx)` | Get declined requirements |
| `client.PhoneNumbersAPI.DeletePhoneNumberStockWatch(ctx)` | Stop watching a country |
| `client.PhoneNumbersAPI.CancelPhoneNumberPortIn(ctx)` | Cancel a port-in |
| `client.PhoneNumbersAPI.CheckPhoneNumberAvailability(ctx)` | Check country availability |
| `client.PhoneNumbersAPI.CheckPhoneNumberPortability(ctx)` | Check portability |
| `client.PhoneNumbersAPI.PurchasePhoneNumber(ctx)` | Purchase phone number |
| `client.PhoneNumbersAPI.ReleasePhoneNumber(ctx)` | Release phone number |
| `client.PhoneNumbersAPI.RemediatePhoneNumber(ctx)` | Resubmit a declined number |
| `client.PhoneNumbersAPI.ReplyToPhoneNumberReviewer(ctx)` | Reply to the regulatory reviewer |
| `client.PhoneNumbersAPI.RequestPhoneNumberWhatsAppCode(ctx)` | Request the WhatsApp verification code for a number |
| `client.PhoneNumbersAPI.RespondToPhoneNumberReviewer(ctx)` | Respond to the regulatory reviewer (message + corrections) |
| `client.PhoneNumbersAPI.ReviewPhoneNumberKycPacket(ctx)` | Pre-review a KYC packet |
| `client.PhoneNumbersAPI.SearchAvailablePhoneNumbers(ctx)` | Search available numbers |
| `client.PhoneNumbersAPI.SubmitPhoneNumberKyc(ctx)` | Submit KYC |
| `client.PhoneNumbersAPI.UploadPhoneNumberKycDocument(ctx)` | Upload a KYC document |
| `client.PhoneNumbersAPI.UploadPhoneNumberPortInDocument(ctx)` | Upload a porting document |
| `client.PhoneNumbersAPI.ValidatePhoneNumberKycAddress(ctx)` | Pre-validate KYC address |
| `client.PhoneNumbersAPI.ViewPhoneNumberKycDocument(ctx)` | View a KYC document on file |

### Product Catalogs
| Method | Description |
|--------|-------------|
| `client.ProductCatalogsAPI.ListAdCatalogFeedUploads(ctx)` | List a feed's uploads |
| `client.ProductCatalogsAPI.ListAdCatalogFeeds(ctx)` | List a catalog's product feeds |
| `client.ProductCatalogsAPI.ListAdCatalogProductSets(ctx)` | List a catalog's product sets |
| `client.ProductCatalogsAPI.ListAdCatalogProducts(ctx)` | List a catalog's products |
| `client.ProductCatalogsAPI.ListAdCatalogs(ctx)` | List Meta product catalogs |
| `client.ProductCatalogsAPI.CreateAdCatalog(ctx)` | Create a Meta product catalog |
| `client.ProductCatalogsAPI.CreateAdCatalogFeed(ctx)` | Create a product feed |
| `client.ProductCatalogsAPI.CreateAdCatalogFeedUpload(ctx)` | Fetch a feed file now |
| `client.ProductCatalogsAPI.CreateAdCatalogProduct(ctx)` | Add a product to a catalog |
| `client.ProductCatalogsAPI.CreateAdCatalogProductSet(ctx)` | Create a product set |
| `client.ProductCatalogsAPI.GetAdCatalog(ctx)` | Get a product catalog |
| `client.ProductCatalogsAPI.GetAdCatalogBatch(ctx)` | Get a bulk request's status |
| `client.ProductCatalogsAPI.GetAdCatalogProduct(ctx)` | Get a product |
| `client.ProductCatalogsAPI.UpdateAdCatalogProduct(ctx)` | Update a product |
| `client.ProductCatalogsAPI.UpdateAdCatalogProductSet(ctx)` | Update a product set |
| `client.ProductCatalogsAPI.DeleteAdCatalog(ctx)` | Delete a product catalog |
| `client.ProductCatalogsAPI.DeleteAdCatalogProduct(ctx)` | Delete a product |
| `client.ProductCatalogsAPI.DeleteAdCatalogProductSet(ctx)` | Delete a product set |
| `client.ProductCatalogsAPI.BatchAdCatalogProducts(ctx)` | Create, update or delete products in bulk |

### Products
| Method | Description |
|--------|-------------|
| `client.ProductsAPI.ListProducts(ctx)` | List products |
| `client.ProductsAPI.GetProduct(ctx)` | Get a product |
| `client.ProductsAPI.UpdateProduct(ctx)` | Update a product |

### RCS
| Method | Description |
|--------|-------------|
| `client.RCSAPI.ListRcsAgents(ctx)` | List RCS agents |
| `client.RCSAPI.ListRcsBrands(ctx)` | List RCS brands |
| `client.RCSAPI.ListRcsTestDevices(ctx)` | List RCS test phones |
| `client.RCSAPI.CreateRcsAgent(ctx)` | Request an RCS agent |
| `client.RCSAPI.GetRcsAgent(ctx)` | Get an RCS agent |
| `client.RCSAPI.GetRcsCapabilities(ctx)` | Check RCS capability |
| `client.RCSAPI.UpdateRcsAgent(ctx)` | Update an RCS agent |
| `client.RCSAPI.AddRcsTestDevice(ctx)` | Invite an RCS test phone |
| `client.RCSAPI.DeactivateRcsAgent(ctx)` | Deactivate an RCS agent |
| `client.RCSAPI.RemoveRcsTestDevice(ctx)` | Remove an RCS test phone |
| `client.RCSAPI.RequestRcsAgentLaunch(ctx)` | Send the launch filing |
| `client.RCSAPI.SendRcsMessage(ctx)` | Send an RCS message |
| `client.RCSAPI.UploadRcsAsset(ctx)` | Upload an RCS logo or banner |

### Reach and Frequency
| Method | Description |
|--------|-------------|
| `client.ReachAndFrequencyAPI.CreateRfPrediction(ctx)` | Create reach-frequency prediction |
| `client.ReachAndFrequencyAPI.GetRfPrediction(ctx)` | Get reach-frequency prediction |
| `client.ReachAndFrequencyAPI.CancelRfReservation(ctx)` | Cancel reach-frequency booking |
| `client.ReachAndFrequencyAPI.ReserveRfPrediction(ctx)` | Reserve reach-frequency inventory |

### Reviews (Inbox)
| Method | Description |
|--------|-------------|
| `client.ReviewsAPI.ListInboxReviews(ctx)` | List reviews |
| `client.ReviewsAPI.DeleteInboxReviewReply(ctx)` | Delete review reply |
| `client.ReviewsAPI.ReplyToInboxReview(ctx)` | Reply to review |

### SMS
| Method | Description |
|--------|-------------|
| `client.SMSAPI.ListSmsOptOuts(ctx)` | List SMS opt-outs |
| `client.SMSAPI.ListSmsRegistrations(ctx)` | List carrier registrations |
| `client.SMSAPI.ListSmsSenderIds(ctx)` | List alphanumeric sender IDs |
| `client.SMSAPI.CreateSmsSenderId(ctx)` | Create an alphanumeric sender ID |
| `client.SMSAPI.GetSmsRegistration(ctx)` | Get a carrier registration |
| `client.SMSAPI.DeleteSmsSenderId(ctx)` | Delete an alphanumeric sender ID |
| `client.SMSAPI.AppealSmsRegistration(ctx)` | Appeal a rejected campaign |
| `client.SMSAPI.DeactivateSmsRegistration(ctx)` | Deactivate a brand/campaign registration |
| `client.SMSAPI.DisableSmsOnNumber(ctx)` | Disable SMS on a number |
| `client.SMSAPI.EnableSmsOnNumber(ctx)` | Enable SMS on a number |
| `client.SMSAPI.LookupSmsNumber(ctx)` | Look up carrier + line type |
| `client.SMSAPI.PreflightSmsRegistration(ctx)` | Pre-check a carrier registration |
| `client.SMSAPI.RequestSmsSenderIdLimitIncrease(ctx)` | Request a higher sender ID daily limit |
| `client.SMSAPI.ResendSmsRegistrationOtp(ctx)` | Re-send the sole-prop OTP |
| `client.SMSAPI.RespondToSmsRegistrationReview(ctx)` | Reply to a change request |
| `client.SMSAPI.ReuseSmsRegistrationForNumber(ctx)` | Add number to SMS registration |
| `client.SMSAPI.SendSms(ctx)` | Send an SMS/MMS |
| `client.SMSAPI.ShareSmsRegistration(ctx)` | Create a registration share link |
| `client.SMSAPI.StartSmsRegistration(ctx)` | Start a carrier registration |
| `client.SMSAPI.UploadSmsOptInProof(ctx)` | Upload opt-in form proof for an appeal |
| `client.SMSAPI.UploadSmsOptInProofFile(ctx)` | Upload opt-in form proof |
| `client.SMSAPI.VerifySmsRegistrationOtp(ctx)` | Submit the sole-prop OTP |

### Sequences
| Method | Description |
|--------|-------------|
| `client.SequencesAPI.ListSequenceEnrollments(ctx)` | List enrollments for a sequence |
| `client.SequencesAPI.ListSequences(ctx)` | List sequences |
| `client.SequencesAPI.CreateSequence(ctx)` | Create sequence |
| `client.SequencesAPI.GetSequence(ctx)` | Get sequence with steps |
| `client.SequencesAPI.UpdateSequence(ctx)` | Update sequence |
| `client.SequencesAPI.DeleteSequence(ctx)` | Delete sequence |
| `client.SequencesAPI.ActivateSequence(ctx)` | Activate sequence |
| `client.SequencesAPI.EnrollContacts(ctx)` | Enroll contacts in a sequence |
| `client.SequencesAPI.PauseSequence(ctx)` | Pause sequence |
| `client.SequencesAPI.UnenrollContact(ctx)` | Unenroll contact |

### Slack
| Method | Description |
|--------|-------------|
| `client.SlackAPI.ListSlackMembers(ctx)` | List Slack workspace members |

### Tracking Tags
| Method | Description |
|--------|-------------|
| `client.TrackingTagsAPI.ListTrackingTagEvents(ctx)` | List conversion events |
| `client.TrackingTagsAPI.ListTrackingTagPartners(ctx)` | List partner businesses of a tag |
| `client.TrackingTagsAPI.ListTrackingTagSharedAccounts(ctx)` | List accounts it is shared with |
| `client.TrackingTagsAPI.ListTrackingTagUsers(ctx)` | List tag users |
| `client.TrackingTagsAPI.ListTrackingTags(ctx)` | List tracking tags |
| `client.TrackingTagsAPI.CreateTrackingTag(ctx)` | Create a tracking tag |
| `client.TrackingTagsAPI.CreateTrackingTagEvent(ctx)` | Create a conversion event |
| `client.TrackingTagsAPI.GetAdTrackingTags(ctx)` | Get ad tracking tags |
| `client.TrackingTagsAPI.GetTrackingTag(ctx)` | Get a tracking tag |
| `client.TrackingTagsAPI.GetTrackingTagDiagnostics(ctx)` | Get tag diagnostics |
| `client.TrackingTagsAPI.GetTrackingTagStats(ctx)` | Get aggregated event stats |
| `client.TrackingTagsAPI.GetTrackingTagStoreInstall(ctx)` | Get store install status |
| `client.TrackingTagsAPI.UpdateAdTrackingTags(ctx)` | Set ad tracking tags |
| `client.TrackingTagsAPI.UpdateTrackingTag(ctx)` | Update a tracking tag |
| `client.TrackingTagsAPI.UpdateTrackingTagEvent(ctx)` | Update a conversion event |
| `client.TrackingTagsAPI.DeleteTrackingTagEvent(ctx)` | Delete a conversion event |
| `client.TrackingTagsAPI.AddTrackingTagSharedAccount(ctx)` | Share with an ad account |
| `client.TrackingTagsAPI.AssignTrackingTagUser(ctx)` | Assign a user to a tag |
| `client.TrackingTagsAPI.InstallTrackingTagOnStore(ctx)` | Install on a Shopify store or WordPress site |
| `client.TrackingTagsAPI.RemoveTrackingTagFromStore(ctx)` | Remove from a Shopify store or WordPress site |
| `client.TrackingTagsAPI.RemoveTrackingTagSharedAccount(ctx)` | Stop sharing with an account |
| `client.TrackingTagsAPI.RemoveTrackingTagUser(ctx)` | Remove a user from a tag |

### Twitter Engagement
| Method | Description |
|--------|-------------|
| `client.TwitterEngagementAPI.GetTweet(ctx)` | Look up a tweet |
| `client.TwitterEngagementAPI.BookmarkPost(ctx)` | Bookmark a tweet |
| `client.TwitterEngagementAPI.FollowUser(ctx)` | Follow a user |
| `client.TwitterEngagementAPI.RemoveBookmark(ctx)` | Remove bookmark |
| `client.TwitterEngagementAPI.RetweetPost(ctx)` | Retweet a post |
| `client.TwitterEngagementAPI.SearchTweets(ctx)` | Search recent tweets |
| `client.TwitterEngagementAPI.UndoRetweet(ctx)` | Undo retweet |
| `client.TwitterEngagementAPI.UnfollowUser(ctx)` | Unfollow a user |

### Validate
| Method | Description |
|--------|-------------|
| `client.ValidateAPI.ValidateMedia(ctx)` | Validate media URL |
| `client.ValidateAPI.ValidatePost(ctx)` | Validate post content |
| `client.ValidateAPI.ValidatePostLength(ctx)` | Validate character count |
| `client.ValidateAPI.ValidateSubreddit(ctx)` | Check subreddit existence |

### Verify
| Method | Description |
|--------|-------------|
| `client.VerifyAPI.CreateVerification(ctx)` | Send a verification code |
| `client.VerifyAPI.GetVerification(ctx)` | Get a verification |
| `client.VerifyAPI.CheckVerification(ctx)` | Check a verification code |

### Voice
| Method | Description |
|--------|-------------|
| `client.VoiceAPI.ListSipTrunks(ctx)` | List SIP trunks |
| `client.VoiceAPI.ListVoiceCalls(ctx)` | List phone calls |
| `client.VoiceAPI.CreateSipTrunk(ctx)` | Create a SIP trunk |
| `client.VoiceAPI.CreateVoiceCall(ctx)` | Place an outbound phone call |
| `client.VoiceAPI.CreateVoiceWebSession(ctx)` | Mint a browser softphone session |
| `client.VoiceAPI.GetSipTrunk(ctx)` | Get a SIP trunk |
| `client.VoiceAPI.GetVoiceCall(ctx)` | Get a phone call |
| `client.VoiceAPI.GetVoiceCallEstimate(ctx)` | Estimate call cost |
| `client.VoiceAPI.GetVoiceCallRecording(ctx)` | Get a call recording |
| `client.VoiceAPI.DeleteSipTrunk(ctx)` | Delete a SIP trunk |
| `client.VoiceAPI.AttachNumberToSipTrunk(ctx)` | Attach a number to a SIP trunk |
| `client.VoiceAPI.DetachNumberFromSipTrunk(ctx)` | Detach a number from its SIP trunk |
| `client.VoiceAPI.DialVoiceWebCall(ctx)` | Dial from the browser softphone |
| `client.VoiceAPI.DisableVoiceOnNumber(ctx)` | Disable phone calling on a number |
| `client.VoiceAPI.EnableVoiceOnNumber(ctx)` | Enable phone calling on a number |
| `client.VoiceAPI.EndVoiceCall(ctx)` | Hang up a live call |
| `client.VoiceAPI.RotateSipTrunkCredentials(ctx)` | Rotate a SIP trunk's password |
| `client.VoiceAPI.TransferVoiceCall(ctx)` | Blind-transfer a live call |

### WhatsApp
| Method | Description |
|--------|-------------|
| `client.WhatsAppAPI.ListWhatsAppAccountEvents(ctx)` | List account notifications |
| `client.WhatsAppAPI.ListWhatsAppCatalogs(ctx)` | List the catalogs linked to a WhatsApp number |
| `client.WhatsAppAPI.ListWhatsAppConversions(ctx)` | List conversion events |
| `client.WhatsAppAPI.ListWhatsAppGroupChats(ctx)` | List active groups |
| `client.WhatsAppAPI.ListWhatsAppGroupJoinRequests(ctx)` | List join requests |
| `client.WhatsAppAPI.CreateWhatsAppDataset(ctx)` | Provision CTWA dataset |
| `client.WhatsAppAPI.CreateWhatsAppGroupChat(ctx)` | Create group |
| `client.WhatsAppAPI.CreateWhatsAppGroupInviteLink(ctx)` | Create invite link |
| `client.WhatsAppAPI.CreateWhatsAppTemplate(ctx)` | Create template |
| `client.WhatsAppAPI.GetWhatsAppBlockStatus(ctx)` | Check if a user is blocked |
| `client.WhatsAppAPI.GetWhatsAppBlockedUsers(ctx)` | List blocked users |
| `client.WhatsAppAPI.GetWhatsAppBusinessProfile(ctx)` | Get business profile |
| `client.WhatsAppAPI.GetWhatsAppCommerceSettings(ctx)` | Get a number's commerce settings |
| `client.WhatsAppAPI.GetWhatsAppDataset(ctx)` | Get CTWA conversions dataset |
| `client.WhatsAppAPI.GetWhatsAppDisplayName(ctx)` | Get display name status |
| `client.WhatsAppAPI.GetWhatsAppGroupChat(ctx)` | Get group info |
| `client.WhatsAppAPI.GetWhatsAppMedia(ctx)` | Download WhatsApp media |
| `client.WhatsAppAPI.GetWhatsAppTemplate(ctx)` | Get template |
| `client.WhatsAppAPI.GetWhatsAppTemplateById(ctx)` | Get template by id |
| `client.WhatsAppAPI.GetWhatsAppTemplates(ctx)` | List templates |
| `client.WhatsAppAPI.GetWhatsappBusinessUsername(ctx)` | Get business username |
| `client.WhatsAppAPI.GetWhatsappBusinessUsernameSuggestions(ctx)` | Get username suggestions |
| `client.WhatsAppAPI.UpdateWhatsAppBusinessProfile(ctx)` | Update business profile |
| `client.WhatsAppAPI.UpdateWhatsAppCommerceSettings(ctx)` | Update a number's commerce settings |
| `client.WhatsAppAPI.UpdateWhatsAppDisplayName(ctx)` | Request display name change |
| `client.WhatsAppAPI.UpdateWhatsAppGroupChat(ctx)` | Update group settings |
| `client.WhatsAppAPI.UpdateWhatsAppTemplate(ctx)` | Update template |
| `client.WhatsAppAPI.UpdateWhatsAppTemplateById(ctx)` | Update template by id |
| `client.WhatsAppAPI.DeleteWhatsAppGroupChat(ctx)` | Delete group |
| `client.WhatsAppAPI.DeleteWhatsAppTemplate(ctx)` | Delete template |
| `client.WhatsAppAPI.DeleteWhatsAppTemplateById(ctx)` | Delete template by id |
| `client.WhatsAppAPI.DeleteWhatsappBusinessUsername(ctx)` | Delete business username |
| `client.WhatsAppAPI.AddWhatsAppGroupParticipants(ctx)` | Add participants |
| `client.WhatsAppAPI.ApproveWhatsAppGroupJoinRequests(ctx)` | Approve join requests |
| `client.WhatsAppAPI.BlockWhatsAppUsers(ctx)` | Block users |
| `client.WhatsAppAPI.LinkWhatsAppCatalog(ctx)` | Link a catalog to a WhatsApp number |
| `client.WhatsAppAPI.RegisterWhatsAppNumber(ctx)` | Register a connected WhatsApp number on the Cloud API |
| `client.WhatsAppAPI.RejectWhatsAppGroupJoinRequests(ctx)` | Reject join requests |
| `client.WhatsAppAPI.RemoveWhatsAppGroupParticipants(ctx)` | Remove participants |
| `client.WhatsAppAPI.RequestWhatsAppVerificationCode(ctx)` | Request a Meta re-verification code for a BYO WhatsApp number |
| `client.WhatsAppAPI.SendWhatsAppConversion(ctx)` | Send WhatsApp conversion event |
| `client.WhatsAppAPI.SetWhatsappBusinessUsername(ctx)` | Set business username |
| `client.WhatsAppAPI.UnblockWhatsAppUsers(ctx)` | Unblock users |
| `client.WhatsAppAPI.UnlinkWhatsAppCatalog(ctx)` | Unlink a catalog from a WhatsApp number |
| `client.WhatsAppAPI.UploadWhatsAppProfilePhoto(ctx)` | Upload profile picture |
| `client.WhatsAppAPI.VerifyWhatsAppNumber(ctx)` | Verify the Meta re-verification code for a BYO WhatsApp number |

### WhatsApp Calling
| Method | Description |
|--------|-------------|
| `client.WhatsAppCallingAPI.ListWhatsAppCalls(ctx)` | List call history for an account |
| `client.WhatsAppCallingAPI.GetWhatsAppCall(ctx)` | Get a single call |
| `client.WhatsAppCallingAPI.GetWhatsAppCallEstimate(ctx)` | Estimate per-minute cost |
| `client.WhatsAppCallingAPI.GetWhatsAppCallPermissions(ctx)` | Check call permission |
| `client.WhatsAppCallingAPI.GetWhatsAppCallRecording(ctx)` | Get a call recording |
| `client.WhatsAppCallingAPI.GetWhatsAppCalling(ctx)` | Get calling config for a number |
| `client.WhatsAppCallingAPI.GetWhatsAppCallingConfig(ctx)` | Get calling config for an account |
| `client.WhatsAppCallingAPI.UpdateWhatsAppCalling(ctx)` | Update calling config |
| `client.WhatsAppCallingAPI.UpdateWhatsAppCallingLegacy(ctx)` | Update calling config |
| `client.WhatsAppCallingAPI.DisableWhatsAppCalling(ctx)` | Disable calling on a number |
| `client.WhatsAppCallingAPI.DisableWhatsAppCallingLegacy(ctx)` | Disable calling on a number |
| `client.WhatsAppCallingAPI.EnableWhatsAppCalling(ctx)` | Enable calling on a number |
| `client.WhatsAppCallingAPI.EnableWhatsAppCallingLegacy(ctx)` | Enable calling on a number |
| `client.WhatsAppCallingAPI.InitiateWhatsAppCall(ctx)` | Initiate outbound call |
| `client.WhatsAppCallingAPI.StartWhatsAppCallerIdVerification(ctx)` | Start caller-ID verification for a customer-brought number |
| `client.WhatsAppCallingAPI.VerifyWhatsAppCallerId(ctx)` | Confirm the caller-ID verification code |

### WhatsApp Flows
| Method | Description |
|--------|-------------|
| `client.WhatsAppFlowsAPI.ListWhatsAppFlowResponses(ctx)` | List flow responses |
| `client.WhatsAppFlowsAPI.ListWhatsAppFlowVersions(ctx)` | List flow versions |
| `client.WhatsAppFlowsAPI.ListWhatsAppFlows(ctx)` | List flows |
| `client.WhatsAppFlowsAPI.CreateWhatsAppFlow(ctx)` | Create flow |
| `client.WhatsAppFlowsAPI.GetWhatsAppFlow(ctx)` | Get flow |
| `client.WhatsAppFlowsAPI.GetWhatsAppFlowJson(ctx)` | Get flow JSON asset |
| `client.WhatsAppFlowsAPI.GetWhatsAppFlowPreview(ctx)` | Get flow preview URL |
| `client.WhatsAppFlowsAPI.GetWhatsAppFlowsEncryptionKey(ctx)` | Get Flows encryption key status |
| `client.WhatsAppFlowsAPI.UpdateWhatsAppFlow(ctx)` | Update flow |
| `client.WhatsAppFlowsAPI.DeleteWhatsAppFlow(ctx)` | Delete flow |
| `client.WhatsAppFlowsAPI.DeprecateWhatsAppFlow(ctx)` | Deprecate flow |
| `client.WhatsAppFlowsAPI.PublishWhatsAppFlow(ctx)` | Publish flow |
| `client.WhatsAppFlowsAPI.SendWhatsAppFlowMessage(ctx)` | Send flow message |
| `client.WhatsAppFlowsAPI.SetWhatsAppFlowsEncryptionKey(ctx)` | Register a Flows encryption key |
| `client.WhatsAppFlowsAPI.UploadWhatsAppFlowJson(ctx)` | Upload flow JSON |

### WhatsApp Phone Numbers
| Method | Description |
|--------|-------------|
| `client.WhatsAppPhoneNumbersAPI.ListWhatsAppNumberCountries(ctx)` | List offerable number countries |
| `client.WhatsAppPhoneNumbersAPI.CreateWhatsAppNumberKycLink(ctx)` | Create a hosted KYC link |
| `client.WhatsAppPhoneNumbersAPI.GetWhatsAppNumberInfo(ctx)` | Get number status |
| `client.WhatsAppPhoneNumbersAPI.GetWhatsAppNumberKycForm(ctx)` | Get KYC form spec |
| `client.WhatsAppPhoneNumbersAPI.GetWhatsAppNumberRemediation(ctx)` | Get declined requirements |
| `client.WhatsAppPhoneNumbersAPI.GetWhatsAppPhoneNumber(ctx)` | Get phone number |
| `client.WhatsAppPhoneNumbersAPI.GetWhatsAppPhoneNumbers(ctx)` | List phone numbers |
| `client.WhatsAppPhoneNumbersAPI.GetWhatsAppPricingAnalytics(ctx)` | Get pricing analytics |
| `client.WhatsAppPhoneNumbersAPI.CheckWhatsAppNumberAvailability(ctx)` | Check country availability |
| `client.WhatsAppPhoneNumbersAPI.MoveWhatsAppNumberToProfile(ctx)` | Move a number to another profile |
| `client.WhatsAppPhoneNumbersAPI.PurchaseWhatsAppPhoneNumber(ctx)` | Purchase phone number |
| `client.WhatsAppPhoneNumbersAPI.ReleaseWhatsAppPhoneNumber(ctx)` | Release phone number |
| `client.WhatsAppPhoneNumbersAPI.RemediateWhatsAppNumber(ctx)` | Resubmit a declined number |
| `client.WhatsAppPhoneNumbersAPI.SearchAvailableWhatsAppNumbers(ctx)` | Search available numbers |
| `client.WhatsAppPhoneNumbersAPI.SubmitWhatsAppNumberKyc(ctx)` | Submit KYC |
| `client.WhatsAppPhoneNumbersAPI.UploadWhatsAppNumberKycDocument(ctx)` | Upload a KYC document |
| `client.WhatsAppPhoneNumbersAPI.ValidateWhatsAppNumberKycAddress(ctx)` | Pre-validate KYC address |

### WhatsApp Sandbox
| Method | Description |
|--------|-------------|
| `client.WhatsAppSandboxAPI.ListWhatsAppSandboxSessions(ctx)` | List your sandbox sessions |
| `client.WhatsAppSandboxAPI.CreateWhatsAppSandboxSession(ctx)` | Start a sandbox activation |
| `client.WhatsAppSandboxAPI.DeleteWhatsAppSandboxSession(ctx)` | Revoke a sandbox session |

### WhatsApp Templates
| Method | Description |
|--------|-------------|
| `client.WhatsAppTemplatesAPI.GetWhatsAppLibraryTemplate(ctx)` | Look up a library template |

### Workflows
| Method | Description |
|--------|-------------|
| `client.WorkflowsAPI.ListWorkflowExecutionEvents(ctx)` | Get an execution's timeline |
| `client.WorkflowsAPI.ListWorkflowExecutions(ctx)` | List workflow runs |
| `client.WorkflowsAPI.ListWorkflowVersions(ctx)` | List a workflow's version history |
| `client.WorkflowsAPI.ListWorkflows(ctx)` | List workflows |
| `client.WorkflowsAPI.CreateWorkflow(ctx)` | Create workflow |
| `client.WorkflowsAPI.GetWorkflow(ctx)` | Get workflow with graph |
| `client.WorkflowsAPI.GetWorkflowVersion(ctx)` | Get a specific workflow version |
| `client.WorkflowsAPI.UpdateWorkflow(ctx)` | Update workflow |
| `client.WorkflowsAPI.DeleteWorkflow(ctx)` | Delete workflow |
| `client.WorkflowsAPI.ActivateWorkflow(ctx)` | Activate workflow |
| `client.WorkflowsAPI.DuplicateWorkflow(ctx)` | Duplicate a workflow |
| `client.WorkflowsAPI.PauseWorkflow(ctx)` | Pause workflow |
| `client.WorkflowsAPI.RestoreWorkflowVersion(ctx)` | Restore a workflow version |
| `client.WorkflowsAPI.TriggerWorkflow(ctx)` | Manually start a workflow run |

### iMessage
| Method | Description |
|--------|-------------|
| `client.IMessageAPI.ListImessageAudience(ctx)` | List iMessage audience |
| `client.IMessageAPI.ListImessageAvailableNumbers(ctx)` | List instantly available iMessage numbers |
| `client.IMessageAPI.ListImessageSandboxContacts(ctx)` | List iMessage sandbox contacts |
| `client.IMessageAPI.ListImessageSenderOrders(ctx)` | List iMessage sender orders |
| `client.IMessageAPI.ListImessageSenders(ctx)` | List iMessage senders |
| `client.IMessageAPI.CreateImessageGroup(ctx)` | Start an iMessage group chat |
| `client.IMessageAPI.CreateImessageOptInLink(ctx)` | Create a tracked iMessage opt-in link |
| `client.IMessageAPI.GetImessageGroup(ctx)` | Get an iMessage group |
| `client.IMessageAPI.GetImessageSender(ctx)` | Get iMessage sender status |
| `client.IMessageAPI.UpdateImessageGroup(ctx)` | Rename an iMessage group or change its photo |
| `client.IMessageAPI.UpdateImessageSender(ctx)` | Update an iMessage sender |
| `client.IMessageAPI.AddImessageGroupParticipant(ctx)` | Add a participant to an iMessage group |
| `client.IMessageAPI.AddImessageSandboxContact(ctx)` | Add an iMessage sandbox contact |
| `client.IMessageAPI.CancelImessageSender(ctx)` | Cancel an iMessage sender |
| `client.IMessageAPI.OrderImessageSender(ctx)` | Order a new iMessage sender |
| `client.IMessageAPI.RegisterImessageSender(ctx)` | Register an iMessage sender |
| `client.IMessageAPI.RemoveImessageGroupParticipant(ctx)` | Remove a participant from an iMessage group |
| `client.IMessageAPI.RemoveImessageSandboxContact(ctx)` | Remove an iMessage sandbox contact |
| `client.IMessageAPI.ReserveImessageAvailableNumber(ctx)` | Reserve an available iMessage number |
| `client.IMessageAPI.SetImessageSubscription(ctx)` | Subscribe or opt out an iMessage contact |

### Invites
| Method | Description |
|--------|-------------|
| `client.InvitesAPI.CreateInviteToken(ctx)` | Create invite token |

## SDK Reference

### Posts
| Method | Description |
|--------|-------------|
| `client.PostsAPI.ListPosts(ctx)` | List posts |
| `client.PostsAPI.BulkUploadPosts(ctx)` | Bulk upload from CSV |
| `client.PostsAPI.CreatePost(ctx)` | Create post |
| `client.PostsAPI.GetPost(ctx)` | Get post |
| `client.PostsAPI.UpdatePost(ctx)` | Update post |
| `client.PostsAPI.UpdatePostMetadata(ctx)` | Update post metadata |
| `client.PostsAPI.DeletePost(ctx)` | Delete post |
| `client.PostsAPI.EditPost(ctx)` | Edit published post |
| `client.PostsAPI.RetryPost(ctx)` | Retry failed post |
| `client.PostsAPI.UnpublishPost(ctx)` | Unpublish post |

### Accounts
| Method | Description |
|--------|-------------|
| `client.AccountsAPI.GetAllAccountsHealth(ctx)` | Check accounts health |
| `client.AccountsAPI.ListAccounts(ctx)` | List accounts |
| `client.AccountsAPI.ListBusinessPartners(ctx)` | List partner businesses of the Page |
| `client.AccountsAPI.ListTikTokCommercialMusic(ctx)` | List trending commercial music |
| `client.AccountsAPI.GetAccountHealth(ctx)` | Check account health |
| `client.AccountsAPI.GetAccountPosts(ctx)` | List posts published on the platform |
| `client.AccountsAPI.GetBlueskySettings(ctx)` | Get Bluesky account settings |
| `client.AccountsAPI.GetFollowerStats(ctx)` | Get follower stats |
| `client.GMBReviewsAPI.GetGoogleBusinessReview(ctx)` | Get a review |
| `client.GMBReviewsAPI.GetGoogleBusinessReviews(ctx)` | Get reviews |
| `client.AccountsAPI.GetInstagramFollowStatus(ctx)` | Check whether an Instagram user follows the account |
| `client.LinkedInMentionsAPI.GetLinkedInMentions(ctx)` | Resolve LinkedIn mention |
| `client.AccountsAPI.GetSlackSettings(ctx)` | Get Slack account settings |
| `client.AccountsAPI.GetTikTokCreatorInfo(ctx)` | Get TikTok creator info |
| `client.AccountsAPI.UpdateAccount(ctx)` | Update account |
| `client.AccountsAPI.UpdateBlueskySettings(ctx)` | Update Bluesky account settings |
| `client.AccountsAPI.UpdateSlackSettings(ctx)` | Update Slack account settings |
| `client.AccountsAPI.DeleteAccount(ctx)` | Disconnect account |
| `client.GMBReviewsAPI.DeleteGoogleBusinessReviewReply(ctx)` | Delete a review reply |
| `client.GMBReviewsAPI.BatchGetGoogleBusinessReviews(ctx)` | Batch get reviews |
| `client.AccountsAPI.GrantBusinessPartner(ctx)` | Share the Page with a partner business |
| `client.AccountsAPI.MoveAccountToProfile(ctx)` | Move account to another profile |
| `client.GMBReviewsAPI.ReplyToGoogleBusinessReview(ctx)` | Reply to a review |
| `client.AccountsAPI.RevokeBusinessPartner(ctx)` | Revoke a partner business from the Page |
| `client.AccountsAPI.SearchTikTokLocations(ctx)` | Search TikTok location tags |

### Profiles
| Method | Description |
|--------|-------------|
| `client.ProfilesAPI.ListProfiles(ctx)` | List profiles |
| `client.ProfilesAPI.CreateProfile(ctx)` | Create profile |
| `client.ProfilesAPI.GetProfile(ctx)` | Get profile |
| `client.ProfilesAPI.UpdateProfile(ctx)` | Update profile |
| `client.ProfilesAPI.DeleteProfile(ctx)` | Delete profile |

### Analytics
| Method | Description |
|--------|-------------|
| `client.AnalyticsAPI.GetAnalytics(ctx)` | Get post analytics |
| `client.AnalyticsAPI.GetAnalyticsDashboard(ctx)` | Get an analytics dashboard |
| `client.AnalyticsAPI.GetAnalyticsDelta(ctx)` | Analytics changed since a cursor |
| `client.AnalyticsAPI.GetBestTimeToPost(ctx)` | Get best times to post |
| `client.AnalyticsAPI.GetContentDecay(ctx)` | Get content performance decay |
| `client.AnalyticsAPI.GetDailyMetrics(ctx)` | Get daily aggregated metrics |
| `client.AnalyticsAPI.GetFacebookDemographics(ctx)` | Get Facebook Page demographics |
| `client.AnalyticsAPI.GetFacebookPageInsights(ctx)` | Get Facebook Page insights |
| `client.AnalyticsAPI.GetFacebookPostEarnings(ctx)` | Get Facebook post monetization earnings |
| `client.AnalyticsAPI.GetFacebookPostReactions(ctx)` | Get Facebook post reactions |
| `client.AnalyticsAPI.GetGoogleBusinessPerformance(ctx)` | Get Google Business Profile performance metrics |
| `client.AnalyticsAPI.GetGoogleBusinessSearchKeywords(ctx)` | Get Google Business Profile search keywords |
| `client.AnalyticsAPI.GetInstagramAccountInsights(ctx)` | Get Instagram insights |
| `client.AnalyticsAPI.GetInstagramDemographics(ctx)` | Get Instagram demographics |
| `client.AnalyticsAPI.GetInstagramFollowerHistory(ctx)` | Get Instagram follower history |
| `client.AnalyticsAPI.GetLinkedInAggregateAnalytics(ctx)` | Get LinkedIn aggregate stats |
| `client.AnalyticsAPI.GetLinkedInOrgAggregateAnalytics(ctx)` | Get LinkedIn org analytics |
| `client.AnalyticsAPI.GetLinkedInPostAnalytics(ctx)` | Get LinkedIn post stats |
| `client.AnalyticsAPI.GetLinkedInPostReactions(ctx)` | Get LinkedIn post reactions |
| `client.AnalyticsAPI.GetPostTimeline(ctx)` | Get post analytics timeline |
| `client.AnalyticsAPI.GetPostingFrequency(ctx)` | Get frequency vs engagement |
| `client.AnalyticsAPI.GetTikTokAccountInsights(ctx)` | Get TikTok account-level insights |
| `client.AnalyticsAPI.GetYouTubeChannelInsights(ctx)` | Get YouTube channel insights |
| `client.AnalyticsAPI.GetYouTubeDailyViews(ctx)` | Get YouTube daily views |
| `client.AnalyticsAPI.GetYouTubeDemographics(ctx)` | Get YouTube demographics |
| `client.AnalyticsAPI.GetYouTubeVideoRetention(ctx)` | Get YouTube video retention curve |
| `client.AnalyticsAPI.SyncExternalPosts(ctx)` | Sync an external post |

### Account Groups
| Method | Description |
|--------|-------------|
| `client.AccountGroupsAPI.ListAccountGroups(ctx)` | List groups |
| `client.AccountGroupsAPI.CreateAccountGroup(ctx)` | Create group |
| `client.AccountGroupsAPI.UpdateAccountGroup(ctx)` | Update group |
| `client.AccountGroupsAPI.DeleteAccountGroup(ctx)` | Delete group |

### Queue
| Method | Description |
|--------|-------------|
| `client.QueueAPI.ListQueueSlots(ctx)` | List schedules |
| `client.QueueAPI.CreateQueueSlot(ctx)` | Create schedule |
| `client.QueueAPI.GetNextQueueSlot(ctx)` | Get next available slot |
| `client.QueueAPI.UpdateQueueSlot(ctx)` | Update schedule |
| `client.QueueAPI.DeleteQueueSlot(ctx)` | Delete schedule |
| `client.QueueAPI.PreviewQueue(ctx)` | Preview upcoming slots |

### Webhooks
| Method | Description |
|--------|-------------|
| `client.WebhooksAPI.CreateWebhookSettings(ctx)` | Create webhook |
| `client.WebhooksAPI.GetWebhookLogs(ctx)` | List webhook delivery logs |
| `client.WebhooksAPI.GetWebhookSettings(ctx)` | List webhooks |
| `client.WebhooksAPI.UpdateWebhookSettings(ctx)` | Update webhook |
| `client.WebhooksAPI.DeleteWebhookSettings(ctx)` | Delete webhook |
| `client.WebhooksAPI.RedeliverWebhookEvent(ctx)` | Redeliver a webhook event |
| `client.WebhooksAPI.TestWebhook(ctx)` | Send test webhook |

### API Keys
| Method | Description |
|--------|-------------|
| `client.APIKeysAPI.ListApiKeys(ctx)` | List keys |
| `client.APIKeysAPI.CreateApiKey(ctx)` | Create key |
| `client.APIKeysAPI.DeleteApiKey(ctx)` | Delete key |
| `client.APIKeysAPI.VerifyCredential(ctx)` | Verify credential |

### Media
| Method | Description |
|--------|-------------|
| `client.MediaAPI.GetMediaPresignedUrl(ctx)` | Get upload URL |

### Tools
| Method | Description |
|--------|-------------|
| `client.ToolsAPI.DownloadTikTokVideo(ctx)` | Download a TikTok video |

### Users
| Method | Description |
|--------|-------------|
| `client.UsersAPI.ListUsers(ctx)` | List users |
| `client.UsersAPI.GetUser(ctx)` | Get user |

### Usage
| Method | Description |
|--------|-------------|
| `client.UsageAPI.GetBilling(ctx)` | Account billing snapshot (plan, cycle, balance, caps, status) |
| `client.UsageAPI.GetCallsUsage(ctx)` | Calling usage and cost |
| `client.UsageAPI.GetSmsUsage(ctx)` | SMS usage (volumes) |
| `client.UsageAPI.GetUsage(ctx)` | Usage snapshot (default) or billed-spend metering (with params) |
| `client.UsageAPI.GetUsageStats(ctx)` | Get plan and usage snapshot (plan, limits, payment status) |
| `client.UsageAPI.GetXApiPricing(ctx)` | Get X API pricing table |

### Logs
| Method | Description |
|--------|-------------|
| `client.LogsAPI.ListLogs(ctx)` | List activity logs |

### Connect (OAuth)
| Method | Description |
|--------|-------------|
| `client.ConnectAPI.ListFacebookPages(ctx)` | List Facebook pages |
| `client.ConnectAPI.ListGoogleBusinessLocations(ctx)` | List Google Business Profile locations |
| `client.ConnectAPI.ListInstagramPages(ctx)` | List Pages with a linked Instagram account |
| `client.ConnectAPI.ListLinkedInOrganizations(ctx)` | List LinkedIn orgs |
| `client.ConnectAPI.ListPinterestBoardsForSelection(ctx)` | List Pinterest boards |
| `client.ConnectAPI.ListSlackChannels(ctx)` | List Slack channels for the channel picker |
| `client.ConnectAPI.ListSnapchatProfiles(ctx)` | List Snapchat profiles |
| `client.ConnectAPI.ListWhatsAppPhoneNumbers(ctx)` | List numbers for selection |
| `client.ConnectAPI.CreatePinterestBoard(ctx)` | Create Pinterest board |
| `client.ConnectAPI.CreateYoutubePlaylist(ctx)` | Create YouTube playlist |
| `client.ConnectAPI.GetConnectUrl(ctx)` | Get OAuth connect URL |
| `client.ConnectAPI.GetFacebookPages(ctx)` | List Facebook pages |
| `client.ConnectAPI.GetGmbLocations(ctx)` | List Google Business Profile locations |
| `client.ConnectAPI.GetLinkedInOrganizations(ctx)` | List LinkedIn orgs |
| `client.ConnectAPI.GetPageWebhookSubscription(ctx)` | Read a Facebook Page's webhook subscription |
| `client.ConnectAPI.GetPendingOAuthData(ctx)` | Get pending OAuth data |
| `client.ConnectAPI.GetPinterestBoards(ctx)` | List Pinterest boards |
| `client.ConnectAPI.GetRedditFlairs(ctx)` | List subreddit flairs |
| `client.ConnectAPI.GetRedditSubreddits(ctx)` | List Reddit subreddits |
| `client.ConnectAPI.GetShopifyConnectUrl(ctx)` | Get Shopify OAuth connect URL |
| `client.ConnectAPI.GetSubredditRules(ctx)` | Get subreddit rules |
| `client.ConnectAPI.GetTelegramConnectStatus(ctx)` | Generate Telegram code |
| `client.ConnectAPI.GetWhatsAppSdkConfig(ctx)` | Get Embedded Signup SDK config |
| `client.ConnectAPI.GetWordPressAuthUrl(ctx)` | Get WordPress.com OAuth connect URL |
| `client.ConnectAPI.GetYoutubeCaptions(ctx)` | Get a YouTube video transcript |
| `client.ConnectAPI.GetYoutubePlaylists(ctx)` | List YouTube playlists |
| `client.ConnectAPI.UpdateFacebookPage(ctx)` | Update Facebook page |
| `client.ConnectAPI.UpdateGmbLocation(ctx)` | Update Google Business Profile location |
| `client.ConnectAPI.UpdateLinkedInOrganization(ctx)` | Switch LinkedIn account type |
| `client.ConnectAPI.UpdatePinterestBoards(ctx)` | Set default Pinterest board |
| `client.ConnectAPI.UpdateRedditSubreddits(ctx)` | Set default subreddit |
| `client.ConnectAPI.UpdateYoutubeDefaultPlaylist(ctx)` | Set default YouTube playlist |
| `client.ConnectAPI.AssignGoogleBusinessLocation(ctx)` | Assign Google Business Profile location to another profile |
| `client.ConnectAPI.CompleteMetaAdsBusinessLogin(ctx)` | Complete Meta business login |
| `client.ConnectAPI.CompleteTelegramConnect(ctx)` | Check Telegram status |
| `client.ConnectAPI.CompleteWhatsAppPhoneSelection(ctx)` | Complete number selection |
| `client.ConnectAPI.ConfigureTikTokAdsBrandIdentity(ctx)` | Set TikTok brand identity |
| `client.ConnectAPI.ConnectAds(ctx)` | Connect ads for a platform |
| `client.ConnectAPI.ConnectBlueskyCredentials(ctx)` | Connect Bluesky account |
| `client.ConnectAPI.ConnectDiscordChannel(ctx)` | Connect a Discord channel |
| `client.ConnectAPI.ConnectOpenAIAdsCredentials(ctx)` | Connect an OpenAI Ads account |
| `client.ConnectAPI.ConnectShopifyWithToken(ctx)` | Connect a Shopify store with a custom-app Admin token |
| `client.ConnectAPI.ConnectSlackChannel(ctx)` | Connect a Slack channel |
| `client.ConnectAPI.ConnectWhatsAppCredentials(ctx)` | Connect WhatsApp via credentials |
| `client.ConnectAPI.ConnectWhatsAppEmbeddedSignup(ctx)` | Connect WhatsApp from Embedded Signup |
| `client.ConnectAPI.ConnectWordPressWithApplicationPassword(ctx)` | Connect self-hosted WordPress with an application password |
| `client.ConnectAPI.HandleOAuthCallback(ctx)` | Complete OAuth callback |
| `client.ConnectAPI.InitiateTelegramConnect(ctx)` | Connect Telegram directly |
| `client.ConnectAPI.ResyncPageWebhookSubscription(ctx)` | Re-subscribe a Facebook Page to Zernio's webhooks |
| `client.ConnectAPI.SelectFacebookPage(ctx)` | Select Facebook page |
| `client.ConnectAPI.SelectGoogleBusinessLocation(ctx)` | Select Google Business Profile location |
| `client.ConnectAPI.SelectInstagramAccount(ctx)` | Select the Page whose Instagram account to connect |
| `client.ConnectAPI.SelectLinkedInOrganization(ctx)` | Select LinkedIn org |
| `client.ConnectAPI.SelectPinterestBoard(ctx)` | Select Pinterest board |
| `client.ConnectAPI.SelectSnapchatProfile(ctx)` | Select Snapchat profile |
| `client.ConnectAPI.SetRedditPostFlair(ctx)` | Set Reddit post flair |
| `client.ConnectAPI.VoteRedditThing(ctx)` | Vote on a Reddit post or comment |

### Reddit
| Method | Description |
|--------|-------------|
| `client.RedditSearchAPI.GetRedditFeed(ctx)` | Get subreddit feed |
| `client.RedditSearchAPI.SearchReddit(ctx)` | Search posts |

### Account Settings
| Method | Description |
|--------|-------------|
| `client.AccountSettingsAPI.GetInstagramIceBreakers(ctx)` | Get IG ice breakers |
| `client.AccountSettingsAPI.GetMessengerGetStarted(ctx)` | Get FB Get Started button |
| `client.AccountSettingsAPI.GetMessengerMenu(ctx)` | Get FB persistent menu |
| `client.AccountSettingsAPI.GetTelegramCommands(ctx)` | Get TG bot commands |
| `client.AccountSettingsAPI.DeleteInstagramIceBreakers(ctx)` | Delete IG ice breakers |
| `client.AccountSettingsAPI.DeleteMessengerGetStarted(ctx)` | Delete FB Get Started button |
| `client.AccountSettingsAPI.DeleteMessengerMenu(ctx)` | Delete FB persistent menu |
| `client.AccountSettingsAPI.DeleteTelegramCommands(ctx)` | Delete TG bot commands |
| `client.AccountSettingsAPI.SetInstagramIceBreakers(ctx)` | Set IG ice breakers |
| `client.AccountSettingsAPI.SetMessengerGetStarted(ctx)` | Set FB Get Started button |
| `client.AccountSettingsAPI.SetMessengerMenu(ctx)` | Set FB persistent menu |
| `client.AccountSettingsAPI.SetTelegramCommands(ctx)` | Set TG bot commands |

### Ad Accounts
| Method | Description |
|--------|-------------|
| `client.AdAccountsAPI.ListAccountCallouts(ctx)` | List account callouts |
| `client.AdAccountsAPI.ListAccountSitelinks(ctx)` | List account sitelinks |
| `client.AdAccountsAPI.ListAccountStructuredSnippets(ctx)` | List account snippets |
| `client.AdAccountsAPI.ListAdAccountUsers(ctx)` | Ad account users |
| `client.AdAccountsAPI.ListAdAccounts(ctx)` | List ad accounts |
| `client.AdAccountsAPI.ListAdLabels(ctx)` | List ad labels |
| `client.AdAccountsAPI.ListAdNegativeKeywordLists(ctx)` | List negative keyword lists |
| `client.AdAccountsAPI.ListAdStudies(ctx)` | A/B tests and lift studies |
| `client.AdAccountsAPI.ListAdsBusinessCenters(ctx)` | List TikTok Business Centers |
| `client.AdAccountsAPI.ListAdsInstagramAccounts(ctx)` | List Instagram ad identities |
| `client.AdAccountsAPI.ListAdsInstagramPosts(ctx)` | List Instagram posts to boost |
| `client.AdAccountsAPI.ListAdvertisableApplications(ctx)` | List advertisable apps |
| `client.AdAccountsAPI.ListCustomConversions(ctx)` | List custom conversions |
| `client.AdAccountsAPI.ListHighDemandPeriods(ctx)` | List high-demand periods |
| `client.AdAccountsAPI.ListMetaBusinessUsers(ctx)` | Business users |
| `client.AdAccountsAPI.ListMetaBusinesses(ctx)` | Businesses list |
| `client.AdAccountsAPI.ListPageUsers(ctx)` | Page users of a business |
| `client.AdAccountsAPI.ListTikTokAdPixels(ctx)` | List TikTok ad pixels |
| `client.AdAccountsAPI.ListValueRuleSets(ctx)` | List value rule sets |
| `client.AdAccountsAPI.CreateAdAccount(ctx)` | Create Meta ad account |
| `client.AdAccountsAPI.CreateAdLabel(ctx)` | Create a Google Ads label |
| `client.AdAccountsAPI.CreateAdNegativeKeywordList(ctx)` | Create a negative keyword list |
| `client.AdAccountsAPI.CreateCustomConversion(ctx)` | Create custom conversion |
| `client.AdAccountsAPI.CreateHighDemandPeriod(ctx)` | Schedule a budget increase |
| `client.AdAccountsAPI.CreateValueRuleSet(ctx)` | Create a value rule set |
| `client.AdAccountsAPI.GetAdAccountFinance(ctx)` | Ad account finances |
| `client.AdAccountsAPI.GetAdAccountHierarchy(ctx)` | Get manager account hierarchy |
| `client.AdAccountsAPI.GetAdComments(ctx)` | List comments on an ad |
| `client.AdAccountsAPI.GetAdNegativeKeywordList(ctx)` | Get a negative keyword list |
| `client.AdAccountsAPI.GetAdsActivityLog(ctx)` | Ad account change / audit log |
| `client.AdAccountsAPI.GetDsaDefaults(ctx)` | Get ad account DSA defaults |
| `client.AdAccountsAPI.GetDsaRecommendations(ctx)` | Get DSA recommendations |
| `client.AdAccountsAPI.GetIosFourteenCampaignLimits(ctx)` | Get iOS 14 campaign limits |
| `client.AdAccountsAPI.GetValueRuleSet(ctx)` | Read a value rule set |
| `client.AdAccountsAPI.UpdateAccountCallouts(ctx)` | Update account callouts |
| `client.AdAccountsAPI.UpdateAccountSitelinks(ctx)` | Update account sitelinks |
| `client.AdAccountsAPI.UpdateAccountStructuredSnippets(ctx)` | Update account snippets |
| `client.AdAccountsAPI.UpdateAdAccount(ctx)` | Update ad account settings |
| `client.AdAccountsAPI.UpdateAdAccountManagerLink(ctx)` | Accept, decline, cancel or end a manager link |
| `client.AdAccountsAPI.UpdateAdLabel(ctx)` | Update a Google Ads label |
| `client.AdAccountsAPI.UpdateAdNegativeKeywordList(ctx)` | Rename a negative keyword list |
| `client.AdAccountsAPI.UpdateValueRuleSet(ctx)` | Replace a value rule set |
| `client.AdAccountsAPI.DeleteAdComment(ctx)` | Delete an ad comment |
| `client.AdAccountsAPI.DeleteAdNegativeKeywordList(ctx)` | Delete a negative keyword list |
| `client.AdAccountsAPI.DeleteValueRuleSet(ctx)` | Delete a value rule set |
| `client.AdAccountsAPI.AddAccountCallouts(ctx)` | Add account callouts |
| `client.AdAccountsAPI.AddAccountSitelinks(ctx)` | Add account sitelinks |
| `client.AdAccountsAPI.AddAccountStructuredSnippets(ctx)` | Add account snippets |
| `client.AdAccountsAPI.AssignAdAccountUser(ctx)` | Assign a user to an ad account |
| `client.AdAccountsAPI.AssignPageUser(ctx)` | Assign a user to a Page |
| `client.AdAccountsAPI.AttachAdLabel(ctx)` | Attach a Google Ads label |
| `client.AdAccountsAPI.DetachAdLabel(ctx)` | Detach a Google Ads label |
| `client.AdAccountsAPI.HideAdComment(ctx)` | Hide or unhide an ad comment |
| `client.AdAccountsAPI.InviteAdAccountToManager(ctx)` | Invite a client account to a manager |
| `client.AdAccountsAPI.RemoveAccountCallout(ctx)` | Remove account callout |
| `client.AdAccountsAPI.RemoveAccountSitelink(ctx)` | Remove account sitelink |
| `client.AdAccountsAPI.RemoveAccountStructuredSnippet(ctx)` | Remove account snippet |
| `client.AdAccountsAPI.RemoveAdAccountUser(ctx)` | Remove a user from an ad account |
| `client.AdAccountsAPI.RemoveAdLabel(ctx)` | Remove a Google Ads label |
| `client.AdAccountsAPI.RemovePageUser(ctx)` | Remove a user from a Page |
| `client.AdAccountsAPI.ReplaceAdNegativeKeywordListKeywords(ctx)` | Replace negative list keywords |
| `client.AdAccountsAPI.ReplyToAdComment(ctx)` | Reply to an ad comment |

### Ad Audiences
| Method | Description |
|--------|-------------|
| `client.AdAudiencesAPI.ListAdAudiences(ctx)` | List custom audiences |
| `client.AdAudiencesAPI.CreateAdAudience(ctx)` | Create custom audience |
| `client.AdAudiencesAPI.GetAdAudience(ctx)` | Get audience details |
| `client.AdAudiencesAPI.UpdateAdAudience(ctx)` | Update an audience |
| `client.AdAudiencesAPI.DeleteAdAudience(ctx)` | Delete custom audience |
| `client.AdAudiencesAPI.AddUsersToAdAudience(ctx)` | Add users to audience |
| `client.AdAudiencesAPI.ReplaceAdAudienceCompanies(ctx)` | Replace audience companies |

### Ad Campaigns
| Method | Description |
|--------|-------------|
| `client.AdCampaignsAPI.ListAdCampaigns(ctx)` | List campaigns |
| `client.AdCampaignsAPI.ListAdGroupAssets(ctx)` | List ad-group assets |
| `client.AdCampaignsAPI.ListAdKeywords(ctx)` | List Search keywords |
| `client.AdCampaignsAPI.ListAdSets(ctx)` | List ad sets |
| `client.AdCampaignsAPI.ListAds(ctx)` | List ads |
| `client.AdCampaignsAPI.ListBidStrategies(ctx)` | List portfolio bid strategies |
| `client.AdCampaignsAPI.ListCampaignAssets(ctx)` | List campaign assets |
| `client.AdCampaignsAPI.ListCampaignNegativeKeywordLists(ctx)` | List campaign negative lists |
| `client.AdCampaignsAPI.ListCampaignNegativeKeywords(ctx)` | List campaign-level negative keywords |
| `client.AdCampaignsAPI.ListGoogleAssetGroups(ctx)` | List Performance Max asset groups |
| `client.AdCampaignsAPI.ListGoogleRecommendations(ctx)` | List Google Ads recommendations |
| `client.AdCampaignsAPI.BulkUpdateAdCampaignStatus(ctx)` | Pause or resume many campaigns |
| `client.AdCampaignsAPI.CreateAdCampaign(ctx)` | Create a standalone campaign |
| `client.AdCampaignsAPI.CreateAdSet(ctx)` | Create a standalone ad group |
| `client.AdCampaignsAPI.CreateBidStrategy(ctx)` | Create portfolio bid strategy |
| `client.AdCampaignsAPI.CreateGoogleAssetGroup(ctx)` | Create a Performance Max asset group |
| `client.AdCampaignsAPI.CreateStandaloneAd(ctx)` | Create standalone ad |
| `client.AdCampaignsAPI.GetAd(ctx)` | Get ad details |
| `client.AdCampaignsAPI.GetAdCampaignDetails(ctx)` | Get live campaign details |
| `client.AdCampaignsAPI.GetAdReview(ctx)` | Read the platform's review verdict for an ad |
| `client.AdCampaignsAPI.GetAdSetDetails(ctx)` | Get live ad-set details |
| `client.AdCampaignsAPI.GetAdTree(ctx)` | Get campaign tree |
| `client.AdCampaignsAPI.GetAdsTimeline(ctx)` | Get daily account metrics |
| `client.AdCampaignsAPI.GetCampaignAdSchedule(ctx)` | Read a campaign's ad schedule (dayparting) |
| `client.AdCampaignsAPI.GetCampaignBidding(ctx)` | Read a campaign's current bidding |
| `client.AdCampaignsAPI.GetCampaignConversionGoals(ctx)` | Get campaign conversion goals |
| `client.AdCampaignsAPI.GetCampaignTargeting(ctx)` | Read a Google campaign's device, location, and language targeting |
| `client.AdCampaignsAPI.GetGoogleAssetGroup(ctx)` | Get a Performance Max asset group |
| `client.AdCampaignsAPI.UpdateAd(ctx)` | Update ad |
| `client.AdCampaignsAPI.UpdateAdCampaign(ctx)` | Update a campaign |
| `client.AdCampaignsAPI.UpdateAdCampaignStatus(ctx)` | Pause or resume a campaign |
| `client.AdCampaignsAPI.UpdateAdGroupAssets(ctx)` | Update ad-group assets |
| `client.AdCampaignsAPI.UpdateAdKeyword(ctx)` | Pause or enable a Search keyword |
| `client.AdCampaignsAPI.UpdateAdSet(ctx)` | Update an ad set |
| `client.AdCampaignsAPI.UpdateAdSetStatus(ctx)` | Pause or resume a single ad set |
| `client.AdCampaignsAPI.UpdateAdStatus(ctx)` | Pause or resume a single ad |
| `client.AdCampaignsAPI.UpdateBidStrategy(ctx)` | Update portfolio bid strategy |
| `client.AdCampaignsAPI.UpdateCampaignAdSchedule(ctx)` | Replace a campaign's ad schedule (dayparting) |
| `client.AdCampaignsAPI.UpdateCampaignAssets(ctx)` | Update campaign assets |
| `client.AdCampaignsAPI.UpdateCampaignConversionGoals(ctx)` | Update campaign conversion goals |
| `client.AdCampaignsAPI.UpdateCampaignTargeting(ctx)` | Edit a Google campaign's device, location, or language targeting |
| `client.AdCampaignsAPI.UpdateGoogleAssetGroup(ctx)` | Update a Performance Max asset group |
| `client.AdCampaignsAPI.DeleteAd(ctx)` | Cancel an ad |
| `client.AdCampaignsAPI.DeleteAdCampaign(ctx)` | Delete a campaign |
| `client.AdCampaignsAPI.DeleteAdSet(ctx)` | Delete an ad set |
| `client.AdCampaignsAPI.AddAdKeywords(ctx)` | Add Search ad-group keywords |
| `client.AdCampaignsAPI.ApplyGoogleRecommendations(ctx)` | Apply Google Ads recommendations |
| `client.AdCampaignsAPI.AttachAdGroupAssets(ctx)` | Attach ad-group assets |
| `client.AdCampaignsAPI.AttachCampaignAssets(ctx)` | Attach campaign assets |
| `client.AdCampaignsAPI.BoostPost(ctx)` | Boost post as ad |
| `client.AdCampaignsAPI.DismissGoogleRecommendations(ctx)` | Dismiss Google Ads recommendations |
| `client.AdCampaignsAPI.DuplicateAd(ctx)` | Duplicate an ad |
| `client.AdCampaignsAPI.DuplicateAdCampaign(ctx)` | Duplicate a campaign |
| `client.AdCampaignsAPI.DuplicateAdSet(ctx)` | Duplicate an ad set |
| `client.AdCampaignsAPI.EditGoogleAssetGroupAssets(ctx)` | Link or unlink asset group assets |
| `client.AdCampaignsAPI.RemoveAdGroupAssets(ctx)` | Remove ad-group assets |
| `client.AdCampaignsAPI.RemoveAdKeyword(ctx)` | Remove a Search keyword |
| `client.AdCampaignsAPI.RemoveCampaignAssets(ctx)` | Remove campaign assets |
| `client.AdCampaignsAPI.RemoveGoogleAssetGroup(ctx)` | Remove a Performance Max asset group |
| `client.AdCampaignsAPI.ReplaceCampaignNegativeKeywordLists(ctx)` | Replace campaign negative lists |
| `client.AdCampaignsAPI.ReplaceCampaignNegativeKeywords(ctx)` | Replace campaign-level negative keywords |
| `client.AdCampaignsAPI.ReplaceGoogleListingGroupFilters(ctx)` | Replace an asset group's listing-group tree |

### Ad Creatives
| Method | Description |
|--------|-------------|
| `client.AdCreativesAPI.ListAdCreatives(ctx)` | Creative library |
| `client.AdCreativesAPI.ListAdImages(ctx)` | Ad image library |
| `client.AdCreativesAPI.ListAdVideos(ctx)` | Ad video library |
| `client.AdCreativesAPI.ListAdsTikTokIdentities(ctx)` | List TikTok ad identities |
| `client.AdCreativesAPI.ListPartnershipAdContent(ctx)` | List partnership ad content |
| `client.AdCreativesAPI.ListPartnershipAdPermissions(ctx)` | List partnership permissions |
| `client.AdCreativesAPI.CreateAdCreative(ctx)` | Create a standalone creative |
| `client.AdCreativesAPI.GetAdCreative(ctx)` | Creative details |
| `client.AdCreativesAPI.GetAdMedia(ctx)` | Direct video and image URLs for an ad |
| `client.AdCreativesAPI.GetAdPreviews(ctx)` | Render previews of an existing ad |
| `client.AdCreativesAPI.UpdateAdCreative(ctx)` | Rename a creative |
| `client.AdCreativesAPI.DeleteAdCreative(ctx)` | Delete a creative |
| `client.AdCreativesAPI.DeleteAdVideo(ctx)` | Delete an ad video |
| `client.AdCreativesAPI.GenerateAdPreviews(ctx)` | Render pre-create ad previews |
| `client.AdCreativesAPI.SetPartnershipAdPermission(ctx)` | Set partnership permission |
| `client.AdCreativesAPI.UploadAdImage(ctx)` | Upload an ad image from base64 |
| `client.AdCreativesAPI.UploadAdVideo(ctx)` | Upload an ad video |

### Ad Insights
| Method | Description |
|--------|-------------|
| `client.AdInsightsAPI.ListLocalServicesLeadConversations(ctx)` | List lead conversations |
| `client.AdInsightsAPI.ListLocalServicesLeads(ctx)` | Google Local Services Ads leads |
| `client.AdInsightsAPI.CreateAdInsightsReport(ctx)` | Submit async insights report |
| `client.AdInsightsAPI.GetAdAnalytics(ctx)` | Get ad analytics |
| `client.AdInsightsAPI.GetAdInsightsReport(ctx)` | Poll an async insights report run |
| `client.AdInsightsAPI.GetAdsSearchTerms(ctx)` | Google Ads search terms report |
| `client.AdInsightsAPI.GetCampaignAnalytics(ctx)` | Get campaign analytics |
| `client.AdInsightsAPI.GetTikTokSmartPlusMaterialReport(ctx)` | Per-creative performance inside TikTok Smart+ ads |
| `client.AdInsightsAPI.GenerateKeywordHistoricalMetrics(ctx)` | Get historical keyword metrics |
| `client.AdInsightsAPI.GenerateKeywordIdeas(ctx)` | Generate keyword ideas |
| `client.AdInsightsAPI.QueryAdInsights(ctx)` | Flexible live insights query |

### Ad Library
| Method | Description |
|--------|-------------|
| `client.AdLibraryAPI.SearchAdLibrary(ctx)` | Search the public Ad Library |

### Ad Targeting
| Method | Description |
|--------|-------------|
| `client.AdTargetingAPI.GetLinkedInBidPricing(ctx)` | Suggested bid and budget bounds |
| `client.AdTargetingAPI.GetLinkedInSupplyForecast(ctx)` | Forecast ad delivery |
| `client.AdTargetingAPI.EstimateAdReach(ctx)` | Estimate audience reach |
| `client.AdTargetingAPI.SearchAdInterests(ctx)` | Search targeting interests |
| `client.AdTargetingAPI.SearchAdTargeting(ctx)` | Search targeting options |

### Blogs
| Method | Description |
|--------|-------------|
| `client.BlogsAPI.ListBlogArticles(ctx)` | List blog articles |
| `client.BlogsAPI.ListBlogs(ctx)` | List blogs |
| `client.BlogsAPI.CreateBlog(ctx)` | Create a blog |
| `client.BlogsAPI.CreateBlogArticle(ctx)` | Create a blog article |
| `client.BlogsAPI.GetBlog(ctx)` | Get a blog |
| `client.BlogsAPI.GetBlogArticle(ctx)` | Get a blog article |
| `client.BlogsAPI.UpdateBlog(ctx)` | Update a blog |
| `client.BlogsAPI.UpdateBlogArticle(ctx)` | Update a blog article |
| `client.BlogsAPI.DeleteBlog(ctx)` | Delete a blog |
| `client.BlogsAPI.DeleteBlogArticle(ctx)` | Delete a blog article |

### Branded Calling
| Method | Description |
|--------|-------------|
| `client.BrandedCallingAPI.ListBrandedCallingCallReasons(ctx)` | List pre-approved call reasons |
| `client.BrandedCallingAPI.ListBrandedCallingEnterprises(ctx)` | List registered businesses |
| `client.BrandedCallingAPI.ListBrandedCallingIdentities(ctx)` | List caller identities |
| `client.BrandedCallingAPI.ListBrandedCallingIdentityNumbers(ctx)` | List the numbers on a caller identity |
| `client.BrandedCallingAPI.CreateBrandedCallingEnterprise(ctx)` | Register a business for Branded Calling |
| `client.BrandedCallingAPI.CreateBrandedCallingIdentity(ctx)` | Create a caller identity |
| `client.BrandedCallingAPI.GetBrandedCallingEnterprise(ctx)` | Get a registered business |
| `client.BrandedCallingAPI.GetBrandedCallingIdentity(ctx)` | Get a caller identity |
| `client.BrandedCallingAPI.UpdateBrandedCallingIdentity(ctx)` | Edit or resubmit a caller identity |
| `client.BrandedCallingAPI.DeleteBrandedCallingEnterprise(ctx)` | Delete a registered business |
| `client.BrandedCallingAPI.DeleteBrandedCallingIdentity(ctx)` | Delete a caller identity |
| `client.BrandedCallingAPI.AttachBrandedCallingNumbers(ctx)` | Attach numbers to a verified identity |
| `client.BrandedCallingAPI.ConfirmBrandedCallingAuthorizerEmail(ctx)` | Confirm the authorizer's code |
| `client.BrandedCallingAPI.DetachBrandedCallingNumbers(ctx)` | Detach numbers from an identity |
| `client.BrandedCallingAPI.PreflightBrandedCallingIdentity(ctx)` | Dry-run a caller identity before creating it |
| `client.BrandedCallingAPI.ResendBrandedCallingAuthorizerCode(ctx)` | Resend the authorizer's code |
| `client.BrandedCallingAPI.ShareBrandedCallingIdentityForm(ctx)` | Create a caller identity share link |

### Broadcasts
| Method | Description |
|--------|-------------|
| `client.BroadcastsAPI.ListBroadcastRecipients(ctx)` | List broadcast recipients |
| `client.BroadcastsAPI.ListBroadcasts(ctx)` | List broadcasts |
| `client.BroadcastsAPI.CreateBroadcast(ctx)` | Create broadcast draft |
| `client.BroadcastsAPI.GetBroadcast(ctx)` | Get broadcast details |
| `client.BroadcastsAPI.UpdateBroadcast(ctx)` | Update broadcast |
| `client.BroadcastsAPI.DeleteBroadcast(ctx)` | Delete broadcast |
| `client.BroadcastsAPI.AddBroadcastRecipients(ctx)` | Add recipients to a broadcast |
| `client.BroadcastsAPI.CancelBroadcast(ctx)` | Cancel broadcast |
| `client.BroadcastsAPI.ScheduleBroadcast(ctx)` | Schedule broadcast for later |
| `client.BroadcastsAPI.SendBroadcast(ctx)` | Send broadcast now |

### Business Agent
| Method | Description |
|--------|-------------|
| `client.BusinessAgentAPI.ListBusinessAgentAllowlist(ctx)` | List allowlisted consumers |
| `client.BusinessAgentAPI.ListBusinessAgentConnectorTools(ctx)` | List connector tools |
| `client.BusinessAgentAPI.ListBusinessAgentConnectors(ctx)` | List connectors |
| `client.BusinessAgentAPI.ListBusinessAgentFaqs(ctx)` | List FAQs |
| `client.BusinessAgentAPI.ListBusinessAgentFiles(ctx)` | List knowledge files |
| `client.BusinessAgentAPI.ListBusinessAgentSettings(ctx)` | List agent settings |
| `client.BusinessAgentAPI.ListBusinessAgentSkills(ctx)` | List skills |
| `client.BusinessAgentAPI.ListBusinessAgentUiSkills(ctx)` | List UI skills |
| `client.BusinessAgentAPI.ListBusinessAgentWebsites(ctx)` | List crawled websites |
| `client.BusinessAgentAPI.CreateBusinessAgentConnector(ctx)` | Create a connector |
| `client.BusinessAgentAPI.CreateBusinessAgentConnectorTool(ctx)` | Create a connector tool |
| `client.BusinessAgentAPI.CreateBusinessAgentFaq(ctx)` | Create a FAQ |
| `client.BusinessAgentAPI.CreateBusinessAgentSkill(ctx)` | Create a skill |
| `client.BusinessAgentAPI.CreateBusinessAgentUiSkill(ctx)` | Create a UI skill |
| `client.BusinessAgentAPI.GetBusinessAgentBudget(ctx)` | Get usage budgets |
| `client.BusinessAgentAPI.GetBusinessAgentBusinessInformation(ctx)` | Get business information |
| `client.BusinessAgentAPI.GetBusinessAgentConnector(ctx)` | Get a connector |
| `client.BusinessAgentAPI.GetBusinessAgentConnectorLogs(ctx)` | Get connector failure logs |
| `client.BusinessAgentAPI.GetBusinessAgentConnectorTool(ctx)` | Get a connector tool |
| `client.BusinessAgentAPI.GetBusinessAgentEvent(ctx)` | Get a business event status |
| `client.BusinessAgentAPI.GetBusinessAgentFaq(ctx)` | Get a FAQ |
| `client.BusinessAgentAPI.GetBusinessAgentFile(ctx)` | Get a knowledge file |
| `client.BusinessAgentAPI.GetBusinessAgentSkill(ctx)` | Get a skill |
| `client.BusinessAgentAPI.GetBusinessAgentStatus(ctx)` | Get agent setup status |
| `client.BusinessAgentAPI.GetBusinessAgentUiSkill(ctx)` | Get a UI skill |
| `client.BusinessAgentAPI.GetBusinessAgentWebsite(ctx)` | Get a crawled website |
| `client.BusinessAgentAPI.UpdateBusinessAgentConnector(ctx)` | Update a connector |
| `client.BusinessAgentAPI.UpdateBusinessAgentConnectorTool(ctx)` | Update a connector tool |
| `client.BusinessAgentAPI.UpdateBusinessAgentFaq(ctx)` | Update a FAQ |
| `client.BusinessAgentAPI.UpdateBusinessAgentSettings(ctx)` | Update agent settings |
| `client.BusinessAgentAPI.UpdateBusinessAgentSkill(ctx)` | Update a skill |
| `client.BusinessAgentAPI.UpdateBusinessAgentUiSkill(ctx)` | Update a UI skill |
| `client.BusinessAgentAPI.UpdateBusinessAgentWebsite(ctx)` | Update a crawled website |
| `client.BusinessAgentAPI.DeleteBusinessAgentConnector(ctx)` | Delete a connector |
| `client.BusinessAgentAPI.DeleteBusinessAgentConnectorTool(ctx)` | Delete a connector tool |
| `client.BusinessAgentAPI.DeleteBusinessAgentFaq(ctx)` | Delete a FAQ |
| `client.BusinessAgentAPI.DeleteBusinessAgentFile(ctx)` | Delete a knowledge file |
| `client.BusinessAgentAPI.DeleteBusinessAgentSkill(ctx)` | Delete a skill |
| `client.BusinessAgentAPI.DeleteBusinessAgentUiSkill(ctx)` | Delete a UI skill |
| `client.BusinessAgentAPI.DeleteBusinessAgentWebsite(ctx)` | Remove a crawled website |
| `client.BusinessAgentAPI.AddBusinessAgentAllowlistEntry(ctx)` | Allowlist a consumer |
| `client.BusinessAgentAPI.AddBusinessAgentWebsite(ctx)` | Add a website to crawl |
| `client.BusinessAgentAPI.OnboardBusinessAgent(ctx)` | Create the agent |
| `client.BusinessAgentAPI.ReadBusinessAgentEvals(ctx)` | Read evaluation data |
| `client.BusinessAgentAPI.RefreshBusinessAgentConnectorTools(ctx)` | Refresh MCP connector tools |
| `client.BusinessAgentAPI.RemoveBusinessAgentAllowlistEntry(ctx)` | Remove an allowlisted consumer |
| `client.BusinessAgentAPI.ReplaceBusinessAgentBudget(ctx)` | Replace usage budgets |
| `client.BusinessAgentAPI.ReplaceBusinessAgentBusinessInformation(ctx)` | Replace business information |
| `client.BusinessAgentAPI.ResetBusinessAgentBusinessInformation(ctx)` | Reset business information |
| `client.BusinessAgentAPI.RunBusinessAgentConnectorTool(ctx)` | Run a connector tool once |
| `client.BusinessAgentAPI.SendBusinessAgentEvent(ctx)` | Send a business event |
| `client.BusinessAgentAPI.SendBusinessAgentTestMessage(ctx)` | Send a test message |
| `client.BusinessAgentAPI.SetBusinessAgentConnectorCredentials(ctx)` | Set connector credentials |
| `client.BusinessAgentAPI.StartBusinessAgentEvalRun(ctx)` | Start an evaluation run |
| `client.BusinessAgentAPI.UploadBusinessAgentFile(ctx)` | Upload a knowledge file |

### Calls
| Method | Description |
|--------|-------------|
| `client.CallsAPI.ListCalls(ctx)` | List all calls (unified history) |
| `client.CallsAPI.GetCall(ctx)` | Get a call (any channel) |
| `client.CallsAPI.GetCallRecording(ctx)` | Get a call recording |

### Changelog
| Method | Description |
|--------|-------------|
| `client.ChangelogAPI.ListChangelog(ctx)` | List API changelog entries |

### Comment Automations
| Method | Description |
|--------|-------------|
| `client.CommentAutomationsAPI.ListCommentAutomationLogs(ctx)` | List automation logs |
| `client.CommentAutomationsAPI.ListCommentAutomations(ctx)` | List comment-to-DM automations |
| `client.CommentAutomationsAPI.CreateCommentAutomation(ctx)` | Create comment-to-DM automation |
| `client.CommentAutomationsAPI.GetCommentAutomation(ctx)` | Get automation details |
| `client.CommentAutomationsAPI.UpdateCommentAutomation(ctx)` | Update automation settings |
| `client.CommentAutomationsAPI.DeleteCommentAutomation(ctx)` | Delete automation |

### Comments (Inbox)
| Method | Description |
|--------|-------------|
| `client.CommentsAPI.ListInboxComments(ctx)` | List commented posts |
| `client.CommentsAPI.GetInboxPostComments(ctx)` | Get post comments |
| `client.CommentsAPI.DeleteInboxComment(ctx)` | Delete comment |
| `client.CommentsAPI.EditInboxComment(ctx)` | Edit comment |
| `client.CommentsAPI.HideInboxComment(ctx)` | Hide comment |
| `client.CommentsAPI.LikeInboxComment(ctx)` | Like comment |
| `client.CommentsAPI.LikePost(ctx)` | Like post |
| `client.CommentsAPI.PinInboxComment(ctx)` | Pin comment |
| `client.CommentsAPI.ReplyToInboxPost(ctx)` | Reply to comment |
| `client.CommentsAPI.SendPrivateReplyToComment(ctx)` | Send private reply |
| `client.CommentsAPI.SetCommentModeration(ctx)` | Set comment moderation status |
| `client.CommentsAPI.UnhideInboxComment(ctx)` | Unhide comment |
| `client.CommentsAPI.UnlikeInboxComment(ctx)` | Unlike comment |
| `client.CommentsAPI.UnlikePost(ctx)` | Unlike post |
| `client.CommentsAPI.UnpinInboxComment(ctx)` | Unpin comment |

### Commerce
| Method | Description |
|--------|-------------|
| `client.CommerceAPI.ListCommerceCatalogSyncs(ctx)` | List catalog syncs |
| `client.CommerceAPI.ListCommerceChannels(ctx)` | List sales channels |
| `client.CommerceAPI.ListCommerceCollectionMetafields(ctx)` | List collection metafields |
| `client.CommerceAPI.ListCommerceCollections(ctx)` | List collections |
| `client.CommerceAPI.ListCommerceDiscounts(ctx)` | List discounts |
| `client.CommerceAPI.ListCommerceInventory(ctx)` | Get a product's stock |
| `client.CommerceAPI.ListCommerceLocations(ctx)` | List locations |
| `client.CommerceAPI.ListCommerceMarkets(ctx)` | List markets |
| `client.CommerceAPI.ListCommerceMenus(ctx)` | List navigation menus |
| `client.CommerceAPI.ListCommerceMetaobjectDefinitions(ctx)` | List metaobject definitions |
| `client.CommerceAPI.ListCommerceMetaobjects(ctx)` | List metaobjects of a type |
| `client.CommerceAPI.ListCommercePages(ctx)` | List pages |
| `client.CommerceAPI.ListCommercePriceLists(ctx)` | List price lists |
| `client.CommerceAPI.ListCommerceProductMetafields(ctx)` | List product metafields |
| `client.CommerceAPI.ListCommerceProducts(ctx)` | List products |
| `client.CommerceAPI.ListCommerceRedirects(ctx)` | List URL redirects |
| `client.CommerceAPI.CreateCommerceCatalogSync(ctx)` | Sync a store into a Meta catalog |
| `client.CommerceAPI.CreateCommerceCollection(ctx)` | Create a collection |
| `client.CommerceAPI.CreateCommerceDiscount(ctx)` | Create a discount |
| `client.CommerceAPI.CreateCommerceMenu(ctx)` | Create a navigation menu |
| `client.CommerceAPI.CreateCommerceMetaobject(ctx)` | Create a metaobject |
| `client.CommerceAPI.CreateCommercePage(ctx)` | Create a page |
| `client.CommerceAPI.CreateCommerceProduct(ctx)` | Create a product |
| `client.CommerceAPI.CreateCommerceProductOptions(ctx)` | Add options |
| `client.CommerceAPI.CreateCommerceProductVariants(ctx)` | Add variants |
| `client.CommerceAPI.CreateCommerceRedirect(ctx)` | Create a URL redirect |
| `client.CommerceAPI.GetCommerceCatalogSync(ctx)` | Get a catalog sync |
| `client.CommerceAPI.GetCommerceCollection(ctx)` | Get a collection |
| `client.CommerceAPI.GetCommerceDiscount(ctx)` | Get a discount |
| `client.CommerceAPI.GetCommerceMenu(ctx)` | Get a navigation menu |
| `client.CommerceAPI.GetCommerceMetaobject(ctx)` | Get a metaobject |
| `client.CommerceAPI.GetCommercePage(ctx)` | Get a page |
| `client.CommerceAPI.GetCommerceProduct(ctx)` | Get a product |
| `client.CommerceAPI.GetCommerceStore(ctx)` | Get a store |
| `client.CommerceAPI.UpdateCommerceCollection(ctx)` | Update a collection |
| `client.CommerceAPI.UpdateCommerceDiscount(ctx)` | Update a discount |
| `client.CommerceAPI.UpdateCommerceMenu(ctx)` | Replace a navigation menu |
| `client.CommerceAPI.UpdateCommerceMetaobject(ctx)` | Update a metaobject |
| `client.CommerceAPI.UpdateCommercePage(ctx)` | Update a page |
| `client.CommerceAPI.UpdateCommerceProduct(ctx)` | Update a product |
| `client.CommerceAPI.UpdateCommerceProductPrices(ctx)` | Update variant prices |
| `client.CommerceAPI.UpdateCommerceRedirect(ctx)` | Update a URL redirect |
| `client.CommerceAPI.DeleteCommerceCatalogSync(ctx)` | Stop a catalog sync |
| `client.CommerceAPI.DeleteCommerceCollection(ctx)` | Delete a collection |
| `client.CommerceAPI.DeleteCommerceCollectionMetafields(ctx)` | Delete collection metafields |
| `client.CommerceAPI.DeleteCommerceDiscount(ctx)` | Delete a discount |
| `client.CommerceAPI.DeleteCommerceMarketingActivity(ctx)` | Delete a marketing activity |
| `client.CommerceAPI.DeleteCommerceMenu(ctx)` | Delete a navigation menu |
| `client.CommerceAPI.DeleteCommerceMetaobject(ctx)` | Delete a metaobject |
| `client.CommerceAPI.DeleteCommercePage(ctx)` | Delete a page |
| `client.CommerceAPI.DeleteCommercePriceListPrices(ctx)` | Remove fixed prices |
| `client.CommerceAPI.DeleteCommerceProductMetafields(ctx)` | Delete product metafields |
| `client.CommerceAPI.DeleteCommerceProductOptions(ctx)` | Delete options |
| `client.CommerceAPI.DeleteCommerceProductVariants(ctx)` | Delete variants |
| `client.CommerceAPI.DeleteCommerceRedirect(ctx)` | Delete a URL redirect |
| `client.CommerceAPI.AddCommerceDiscountCodes(ctx)` | Add codes to a discount |
| `client.CommerceAPI.AddCommerceMarketingEngagement(ctx)` | Report daily engagement |
| `client.CommerceAPI.AddCommerceProductImages(ctx)` | Add images |
| `client.CommerceAPI.ChangeCommerceCollectionChannels(ctx)` | Publish or unpublish a collection |
| `client.CommerceAPI.ChangeCommerceCollectionProducts(ctx)` | Add or remove products in a collection |
| `client.CommerceAPI.ChangeCommerceInventory(ctx)` | Set or adjust stock |
| `client.CommerceAPI.ChangeCommerceProductChannels(ctx)` | Publish or unpublish a product |
| `client.CommerceAPI.ChangeCommerceProductState(ctx)` | Activate, deactivate, archive or delete products |
| `client.CommerceAPI.ChangeCommerceProductTags(ctx)` | Add or remove tags in bulk |
| `client.CommerceAPI.DuplicateCommerceProduct(ctx)` | Duplicate a product |
| `client.CommerceAPI.RemoveCommerceProductImages(ctx)` | Remove images |
| `client.CommerceAPI.ReorderCommerceCollectionProducts(ctx)` | Reorder products in a collection |
| `client.CommerceAPI.ReorderCommerceProductImages(ctx)` | Reorder images |
| `client.CommerceAPI.RunCommerceCatalogSync(ctx)` | Run a catalog sync now |
| `client.CommerceAPI.SetCommerceCollectionMetafields(ctx)` | Set collection metafields |
| `client.CommerceAPI.SetCommerceDiscountActive(ctx)` | Activate or deactivate a discount |
| `client.CommerceAPI.SetCommercePriceListPrices(ctx)` | Set fixed prices |
| `client.CommerceAPI.SetCommerceProductMetafields(ctx)` | Set product metafields |
| `client.CommerceAPI.UpsertCommerceMarketingActivity(ctx)` | Record a marketing activity |

### Connected Apps
| Method | Description |
|--------|-------------|
| `client.ConnectedAppsAPI.ListConnectedApps(ctx)` | List connected apps |
| `client.ConnectedAppsAPI.RevokeConnectedApp(ctx)` | Revoke connected app |

### Contacts
| Method | Description |
|--------|-------------|
| `client.ContactsAPI.ListContacts(ctx)` | List contacts |
| `client.ContactsAPI.BulkCreateContacts(ctx)` | Bulk create contacts |
| `client.ContactsAPI.CreateContact(ctx)` | Create contact |
| `client.ContactsAPI.GetContact(ctx)` | Get contact |
| `client.ContactsAPI.GetContactChannels(ctx)` | List channels for a contact |
| `client.ContactsAPI.UpdateContact(ctx)` | Update contact |
| `client.ContactsAPI.DeleteContact(ctx)` | Delete contact |

### Conversions
| Method | Description |
|--------|-------------|
| `client.ConversionsAPI.ListAdConversionGoals(ctx)` | List account conversion goals |
| `client.ConversionsAPI.ListConversionActions(ctx)` | List conversion actions |
| `client.ConversionsAPI.ListConversionAssociations(ctx)` | List associated campaigns |
| `client.ConversionsAPI.ListConversionDestinations(ctx)` | List conversion destinations |
| `client.ConversionsAPI.ListCustomConversionGoals(ctx)` | List custom conversion goals |
| `client.ConversionsAPI.CreateConversionAction(ctx)` | Create website conversion action |
| `client.ConversionsAPI.CreateConversionDestination(ctx)` | Create a conversion destination |
| `client.ConversionsAPI.CreateCustomConversionGoal(ctx)` | Create a custom conversion goal |
| `client.ConversionsAPI.GetConversionDestination(ctx)` | Get a conversion destination |
| `client.ConversionsAPI.GetConversionMetrics(ctx)` | Get attribution metrics |
| `client.ConversionsAPI.GetConversionsQuality(ctx)` | Get Event Match Quality |
| `client.ConversionsAPI.UpdateAdConversionGoals(ctx)` | Update account conversion goals |
| `client.ConversionsAPI.UpdateConversionAction(ctx)` | Set a conversion action primary or secondary |
| `client.ConversionsAPI.UpdateConversionDestination(ctx)` | Update a conversion destination |
| `client.ConversionsAPI.UpdateCustomConversionGoal(ctx)` | Update a custom conversion goal |
| `client.ConversionsAPI.DeleteConversionDestination(ctx)` | Delete a conversion destination |
| `client.ConversionsAPI.AddConversionAssociations(ctx)` | Associate campaigns |
| `client.ConversionsAPI.AdjustConversions(ctx)` | Adjust uploaded conversions |
| `client.ConversionsAPI.RemoveConversionAssociations(ctx)` | Remove associated campaigns |
| `client.ConversionsAPI.RemoveCustomConversionGoal(ctx)` | Remove a custom conversion goal |
| `client.ConversionsAPI.SendConversions(ctx)` | Send conversion events |

### Custom Fields
| Method | Description |
|--------|-------------|
| `client.CustomFieldsAPI.ListCustomFields(ctx)` | List custom field definitions |
| `client.CustomFieldsAPI.CreateCustomField(ctx)` | Create custom field |
| `client.CustomFieldsAPI.UpdateCustomField(ctx)` | Update custom field |
| `client.CustomFieldsAPI.DeleteCustomField(ctx)` | Delete custom field |
| `client.CustomFieldsAPI.ClearContactFieldValue(ctx)` | Clear custom field value |
| `client.CustomFieldsAPI.SetContactFieldValue(ctx)` | Set custom field value |

### Discord
| Method | Description |
|--------|-------------|
| `client.DiscordAPI.ListDiscordGuildMembers(ctx)` | List Discord guild members |
| `client.DiscordAPI.ListDiscordGuildRoles(ctx)` | List Discord guild roles |
| `client.DiscordAPI.ListDiscordPinnedMessages(ctx)` | List pinned messages |
| `client.DiscordAPI.ListDiscordScheduledEvents(ctx)` | List Discord scheduled events |
| `client.DiscordAPI.CreateDiscordGuildRole(ctx)` | Create a Discord guild role |
| `client.DiscordAPI.CreateDiscordScheduledEvent(ctx)` | Create a Discord scheduled event |
| `client.DiscordAPI.CreateDiscordThread(ctx)` | Create a Discord public thread |
| `client.DiscordAPI.GetDiscordChannels(ctx)` | List Discord guild channels |
| `client.DiscordAPI.GetDiscordGuildMember(ctx)` | Get a Discord guild member |
| `client.DiscordAPI.GetDiscordScheduledEvent(ctx)` | Get a Discord scheduled event |
| `client.DiscordAPI.GetDiscordSettings(ctx)` | Get Discord account settings |
| `client.DiscordAPI.UpdateDiscordScheduledEvent(ctx)` | Update a Discord scheduled event |
| `client.DiscordAPI.UpdateDiscordSettings(ctx)` | Update Discord settings |
| `client.DiscordAPI.DeleteDiscordGuildRole(ctx)` | Delete a Discord guild role |
| `client.DiscordAPI.DeleteDiscordMessage(ctx)` | Delete a Discord channel message |
| `client.DiscordAPI.DeleteDiscordScheduledEvent(ctx)` | Delete a Discord scheduled event |
| `client.DiscordAPI.AddDiscordMemberRole(ctx)` | Assign a role to a guild member |
| `client.DiscordAPI.CrosspostDiscordMessage(ctx)` | Crosspost Discord message |
| `client.DiscordAPI.EditDiscordGuildRole(ctx)` | Edit a Discord guild role |
| `client.DiscordAPI.PinDiscordMessage(ctx)` | Pin a Discord message |
| `client.DiscordAPI.RemoveDiscordMemberRole(ctx)` | Remove a role from a guild member |
| `client.DiscordAPI.SearchDiscordGuildMembers(ctx)` | Search Discord guild members |
| `client.DiscordAPI.SendDiscordDirectMessage(ctx)` | Send a Discord Direct Message |
| `client.DiscordAPI.UnpinDiscordMessage(ctx)` | Unpin a Discord message |

### Feedback
| Method | Description |
|--------|-------------|
| `client.FeedbackAPI.SubmitFeedback(ctx)` | Submit feedback |

### GMB Attributes
| Method | Description |
|--------|-------------|
| `client.GMBAttributesAPI.GetGmbAttributeMetadata(ctx)` | Get attribute metadata |
| `client.GMBAttributesAPI.GetGoogleBusinessAttributes(ctx)` | Get attributes |
| `client.GMBAttributesAPI.UpdateGoogleBusinessAttributes(ctx)` | Update attributes |

### GMB Food Menus
| Method | Description |
|--------|-------------|
| `client.GMBFoodMenusAPI.GetGoogleBusinessFoodMenus(ctx)` | Get food menus |
| `client.GMBFoodMenusAPI.UpdateGoogleBusinessFoodMenus(ctx)` | Update food menus |

### GMB Location Details
| Method | Description |
|--------|-------------|
| `client.GMBLocationDetailsAPI.GetGoogleBusinessLocationDetails(ctx)` | Get location details |
| `client.GMBLocationDetailsAPI.UpdateGoogleBusinessLocationDetails(ctx)` | Update location details |

### GMB Media
| Method | Description |
|--------|-------------|
| `client.GMBMediaAPI.ListGoogleBusinessMedia(ctx)` | List media |
| `client.GMBMediaAPI.CreateGoogleBusinessMedia(ctx)` | Upload photo |
| `client.GMBMediaAPI.DeleteGoogleBusinessMedia(ctx)` | Delete photo |

### GMB Place Actions
| Method | Description |
|--------|-------------|
| `client.GMBPlaceActionsAPI.ListGoogleBusinessPlaceActions(ctx)` | List action links |
| `client.GMBPlaceActionsAPI.CreateGoogleBusinessPlaceAction(ctx)` | Create action link |
| `client.GMBPlaceActionsAPI.UpdateGoogleBusinessPlaceAction(ctx)` | Update action link |
| `client.GMBPlaceActionsAPI.DeleteGoogleBusinessPlaceAction(ctx)` | Delete action link |

### GMB Services
| Method | Description |
|--------|-------------|
| `client.GMBServicesAPI.GetGoogleBusinessServices(ctx)` | Get services |
| `client.GMBServicesAPI.UpdateGoogleBusinessServices(ctx)` | Replace services |

### GMB Verifications
| Method | Description |
|--------|-------------|
| `client.GMBVerificationsAPI.GetGoogleBusinessVerifications(ctx)` | Get verification state |
| `client.GMBVerificationsAPI.CompleteGoogleBusinessVerification(ctx)` | Complete a verification |
| `client.GMBVerificationsAPI.FetchGoogleBusinessVerificationOptions(ctx)` | Fetch verification options |
| `client.GMBVerificationsAPI.StartGoogleBusinessVerification(ctx)` | Start a verification |

### Inbox Analytics
| Method | Description |
|--------|-------------|
| `client.InboxAnalyticsAPI.ListInboxConversationAnalytics(ctx)` | List conversation analytics |
| `client.InboxAnalyticsAPI.GetInboxConversationAnalytics(ctx)` | Get conversation analytics |
| `client.InboxAnalyticsAPI.GetInboxHeatmap(ctx)` | Get day × hour heatmap |
| `client.InboxAnalyticsAPI.GetInboxResponseTime(ctx)` | Get inbox response-time stats |
| `client.InboxAnalyticsAPI.GetInboxSourceBreakdown(ctx)` | Get inbox source breakdown |
| `client.InboxAnalyticsAPI.GetInboxTopAccounts(ctx)` | Get top accounts by inbox volume |
| `client.InboxAnalyticsAPI.GetInboxVolume(ctx)` | Get inbox messaging volume |

### Instagram
| Method | Description |
|--------|-------------|
| `client.InstagramAPI.ListInstagramStories(ctx)` | List active Instagram stories |
| `client.InstagramAPI.GetInstagramAudio(ctx)` | Get Instagram audio metadata |
| `client.InstagramAPI.GetInstagramBusinessDiscovery(ctx)` | Look up a public Instagram Business account |
| `client.InstagramAPI.GetInstagramPublishingLimit(ctx)` | Get Instagram publishing limit |
| `client.InstagramAPI.GetInstagramStoryInsights(ctx)` | Get Instagram story insights |
| `client.InstagramAPI.SearchInstagramAudio(ctx)` | Search Instagram audio |

### Lead Gen
| Method | Description |
|--------|-------------|
| `client.LeadGenAPI.ListFormLeads(ctx)` | List leads for a single form |
| `client.LeadGenAPI.ListLeadForms(ctx)` | List lead forms |
| `client.LeadGenAPI.ListLeads(ctx)` | List submitted leads |
| `client.LeadGenAPI.CreateLeadForm(ctx)` | Create a lead form |
| `client.LeadGenAPI.CreateTestLead(ctx)` | Create a test lead |
| `client.LeadGenAPI.GetLeadForm(ctx)` | Get a lead form |
| `client.LeadGenAPI.DeleteTestLead(ctx)` | Delete a test lead |
| `client.LeadGenAPI.ArchiveLeadForm(ctx)` | Archive a lead form |

### Mentions
| Method | Description |
|--------|-------------|
| `client.MentionsAPI.ListInboxMentions(ctx)` | List mentions |
| `client.MentionsAPI.ReplyToMention(ctx)` | Reply to a mention |

### Messages (Inbox)
| Method | Description |
|--------|-------------|
| `client.MessagesAPI.ListInboxConversations(ctx)` | List conversations |
| `client.MessagesAPI.CreateInboxConversation(ctx)` | Create conversation |
| `client.MessagesAPI.GetInboxConversation(ctx)` | Get conversation |
| `client.MessagesAPI.GetInboxConversationMessages(ctx)` | List messages |
| `client.MessagesAPI.GetMessageAttachment(ctx)` | Resolve message attachment |
| `client.MessagesAPI.UpdateInboxConversation(ctx)` | Update conversation status |
| `client.MessagesAPI.DeleteInboxMessage(ctx)` | Delete message |
| `client.MessagesAPI.AddMessageReaction(ctx)` | Add reaction |
| `client.MessagesAPI.EditInboxMessage(ctx)` | Edit message |
| `client.MessagesAPI.MarkConversationRead(ctx)` | Mark a conversation as read |
| `client.MessagesAPI.RemoveMessageReaction(ctx)` | Remove reaction |
| `client.MessagesAPI.SearchInboxConversations(ctx)` | Search conversations |
| `client.MessagesAPI.SendInboxMessage(ctx)` | Send message |
| `client.MessagesAPI.SendTypingIndicator(ctx)` | Send typing indicator |
| `client.MessagesAPI.SetConversationThreadControl(ctx)` | Hand a conversation to or from Meta Business Agent |
| `client.MessagesAPI.UploadMediaDirect(ctx)` | Upload media file |

### Messaging Ads
| Method | Description |
|--------|-------------|
| `client.MessagingAdsAPI.CreateCallAd(ctx)` | Create Click-to-Call ad |
| `client.MessagingAdsAPI.CreateCtwaAd(ctx)` | Create CTWA ad (deprecated) |
| `client.MessagingAdsAPI.CreateMessagingAd(ctx)` | Create messaging ad |

### Phone Numbers
| Method | Description |
|--------|-------------|
| `client.PhoneNumbersAPI.ListPhoneNumberCountries(ctx)` | List offerable number countries |
| `client.PhoneNumbersAPI.ListPhoneNumberPortIns(ctx)` | List port-in orders |
| `client.PhoneNumbersAPI.ListPhoneNumberStockWatches(ctx)` | List stock watches |
| `client.PhoneNumbersAPI.ListPhoneNumbers(ctx)` | List phone numbers |
| `client.PhoneNumbersAPI.CreatePhoneNumberKycLink(ctx)` | Create a hosted KYC link |
| `client.PhoneNumbersAPI.CreatePhoneNumberPortIn(ctx)` | Port numbers in |
| `client.PhoneNumbersAPI.CreatePhoneNumberStockWatch(ctx)` | Watch an out-of-stock country |
| `client.PhoneNumbersAPI.GetPhoneNumber(ctx)` | Get phone number |
| `client.PhoneNumbersAPI.GetPhoneNumberClaim(ctx)` | Resolve a number claim |
| `client.PhoneNumbersAPI.GetPhoneNumberKycForm(ctx)` | Get KYC form spec |
| `client.PhoneNumbersAPI.GetPhoneNumberPortClaim(ctx)` | Resolve a port claim |
| `client.PhoneNumbersAPI.GetPhoneNumberPortInOrderRequirements(ctx)` | A port-in order's pending requirements |
| `client.PhoneNumbersAPI.GetPhoneNumberPortInRequirements(ctx)` | Country porting requirements |
| `client.PhoneNumbersAPI.GetPhoneNumberRemediation(ctx)` | Get declined requirements |
| `client.PhoneNumbersAPI.DeletePhoneNumberStockWatch(ctx)` | Stop watching a country |
| `client.PhoneNumbersAPI.CancelPhoneNumberPortIn(ctx)` | Cancel a port-in |
| `client.PhoneNumbersAPI.CheckPhoneNumberAvailability(ctx)` | Check country availability |
| `client.PhoneNumbersAPI.CheckPhoneNumberPortability(ctx)` | Check portability |
| `client.PhoneNumbersAPI.PurchasePhoneNumber(ctx)` | Purchase phone number |
| `client.PhoneNumbersAPI.ReleasePhoneNumber(ctx)` | Release phone number |
| `client.PhoneNumbersAPI.RemediatePhoneNumber(ctx)` | Resubmit a declined number |
| `client.PhoneNumbersAPI.ReplyToPhoneNumberReviewer(ctx)` | Reply to the regulatory reviewer |
| `client.PhoneNumbersAPI.RequestPhoneNumberWhatsAppCode(ctx)` | Request the WhatsApp verification code for a number |
| `client.PhoneNumbersAPI.RespondToPhoneNumberReviewer(ctx)` | Respond to the regulatory reviewer (message + corrections) |
| `client.PhoneNumbersAPI.ReviewPhoneNumberKycPacket(ctx)` | Pre-review a KYC packet |
| `client.PhoneNumbersAPI.SearchAvailablePhoneNumbers(ctx)` | Search available numbers |
| `client.PhoneNumbersAPI.SubmitPhoneNumberKyc(ctx)` | Submit KYC |
| `client.PhoneNumbersAPI.UploadPhoneNumberKycDocument(ctx)` | Upload a KYC document |
| `client.PhoneNumbersAPI.UploadPhoneNumberPortInDocument(ctx)` | Upload a porting document |
| `client.PhoneNumbersAPI.ValidatePhoneNumberKycAddress(ctx)` | Pre-validate KYC address |
| `client.PhoneNumbersAPI.ViewPhoneNumberKycDocument(ctx)` | View a KYC document on file |

### Product Catalogs
| Method | Description |
|--------|-------------|
| `client.ProductCatalogsAPI.ListAdCatalogFeedUploads(ctx)` | List a feed's uploads |
| `client.ProductCatalogsAPI.ListAdCatalogFeeds(ctx)` | List a catalog's product feeds |
| `client.ProductCatalogsAPI.ListAdCatalogProductSets(ctx)` | List a catalog's product sets |
| `client.ProductCatalogsAPI.ListAdCatalogProducts(ctx)` | List a catalog's products |
| `client.ProductCatalogsAPI.ListAdCatalogs(ctx)` | List Meta product catalogs |
| `client.ProductCatalogsAPI.CreateAdCatalog(ctx)` | Create a Meta product catalog |
| `client.ProductCatalogsAPI.CreateAdCatalogFeed(ctx)` | Create a product feed |
| `client.ProductCatalogsAPI.CreateAdCatalogFeedUpload(ctx)` | Fetch a feed file now |
| `client.ProductCatalogsAPI.CreateAdCatalogProduct(ctx)` | Add a product to a catalog |
| `client.ProductCatalogsAPI.CreateAdCatalogProductSet(ctx)` | Create a product set |
| `client.ProductCatalogsAPI.GetAdCatalog(ctx)` | Get a product catalog |
| `client.ProductCatalogsAPI.GetAdCatalogBatch(ctx)` | Get a bulk request's status |
| `client.ProductCatalogsAPI.GetAdCatalogProduct(ctx)` | Get a product |
| `client.ProductCatalogsAPI.UpdateAdCatalogProduct(ctx)` | Update a product |
| `client.ProductCatalogsAPI.UpdateAdCatalogProductSet(ctx)` | Update a product set |
| `client.ProductCatalogsAPI.DeleteAdCatalog(ctx)` | Delete a product catalog |
| `client.ProductCatalogsAPI.DeleteAdCatalogProduct(ctx)` | Delete a product |
| `client.ProductCatalogsAPI.DeleteAdCatalogProductSet(ctx)` | Delete a product set |
| `client.ProductCatalogsAPI.BatchAdCatalogProducts(ctx)` | Create, update or delete products in bulk |

### Products
| Method | Description |
|--------|-------------|
| `client.ProductsAPI.ListProducts(ctx)` | List products |
| `client.ProductsAPI.GetProduct(ctx)` | Get a product |
| `client.ProductsAPI.UpdateProduct(ctx)` | Update a product |

### RCS
| Method | Description |
|--------|-------------|
| `client.RCSAPI.ListRcsAgents(ctx)` | List RCS agents |
| `client.RCSAPI.ListRcsBrands(ctx)` | List RCS brands |
| `client.RCSAPI.ListRcsTestDevices(ctx)` | List RCS test phones |
| `client.RCSAPI.CreateRcsAgent(ctx)` | Request an RCS agent |
| `client.RCSAPI.GetRcsAgent(ctx)` | Get an RCS agent |
| `client.RCSAPI.GetRcsCapabilities(ctx)` | Check RCS capability |
| `client.RCSAPI.UpdateRcsAgent(ctx)` | Update an RCS agent |
| `client.RCSAPI.AddRcsTestDevice(ctx)` | Invite an RCS test phone |
| `client.RCSAPI.DeactivateRcsAgent(ctx)` | Deactivate an RCS agent |
| `client.RCSAPI.RemoveRcsTestDevice(ctx)` | Remove an RCS test phone |
| `client.RCSAPI.RequestRcsAgentLaunch(ctx)` | Send the launch filing |
| `client.RCSAPI.SendRcsMessage(ctx)` | Send an RCS message |
| `client.RCSAPI.UploadRcsAsset(ctx)` | Upload an RCS logo or banner |

### Reach and Frequency
| Method | Description |
|--------|-------------|
| `client.ReachAndFrequencyAPI.CreateRfPrediction(ctx)` | Create reach-frequency prediction |
| `client.ReachAndFrequencyAPI.GetRfPrediction(ctx)` | Get reach-frequency prediction |
| `client.ReachAndFrequencyAPI.CancelRfReservation(ctx)` | Cancel reach-frequency booking |
| `client.ReachAndFrequencyAPI.ReserveRfPrediction(ctx)` | Reserve reach-frequency inventory |

### Reviews (Inbox)
| Method | Description |
|--------|-------------|
| `client.ReviewsAPI.ListInboxReviews(ctx)` | List reviews |
| `client.ReviewsAPI.DeleteInboxReviewReply(ctx)` | Delete review reply |
| `client.ReviewsAPI.ReplyToInboxReview(ctx)` | Reply to review |

### SMS
| Method | Description |
|--------|-------------|
| `client.SMSAPI.ListSmsOptOuts(ctx)` | List SMS opt-outs |
| `client.SMSAPI.ListSmsRegistrations(ctx)` | List carrier registrations |
| `client.SMSAPI.ListSmsSenderIds(ctx)` | List alphanumeric sender IDs |
| `client.SMSAPI.CreateSmsSenderId(ctx)` | Create an alphanumeric sender ID |
| `client.SMSAPI.GetSmsRegistration(ctx)` | Get a carrier registration |
| `client.SMSAPI.DeleteSmsSenderId(ctx)` | Delete an alphanumeric sender ID |
| `client.SMSAPI.AppealSmsRegistration(ctx)` | Appeal a rejected campaign |
| `client.SMSAPI.DeactivateSmsRegistration(ctx)` | Deactivate a brand/campaign registration |
| `client.SMSAPI.DisableSmsOnNumber(ctx)` | Disable SMS on a number |
| `client.SMSAPI.EnableSmsOnNumber(ctx)` | Enable SMS on a number |
| `client.SMSAPI.LookupSmsNumber(ctx)` | Look up carrier + line type |
| `client.SMSAPI.PreflightSmsRegistration(ctx)` | Pre-check a carrier registration |
| `client.SMSAPI.RequestSmsSenderIdLimitIncrease(ctx)` | Request a higher sender ID daily limit |
| `client.SMSAPI.ResendSmsRegistrationOtp(ctx)` | Re-send the sole-prop OTP |
| `client.SMSAPI.RespondToSmsRegistrationReview(ctx)` | Reply to a change request |
| `client.SMSAPI.ReuseSmsRegistrationForNumber(ctx)` | Add number to SMS registration |
| `client.SMSAPI.SendSms(ctx)` | Send an SMS/MMS |
| `client.SMSAPI.ShareSmsRegistration(ctx)` | Create a registration share link |
| `client.SMSAPI.StartSmsRegistration(ctx)` | Start a carrier registration |
| `client.SMSAPI.UploadSmsOptInProof(ctx)` | Upload opt-in form proof for an appeal |
| `client.SMSAPI.UploadSmsOptInProofFile(ctx)` | Upload opt-in form proof |
| `client.SMSAPI.VerifySmsRegistrationOtp(ctx)` | Submit the sole-prop OTP |

### Sequences
| Method | Description |
|--------|-------------|
| `client.SequencesAPI.ListSequenceEnrollments(ctx)` | List enrollments for a sequence |
| `client.SequencesAPI.ListSequences(ctx)` | List sequences |
| `client.SequencesAPI.CreateSequence(ctx)` | Create sequence |
| `client.SequencesAPI.GetSequence(ctx)` | Get sequence with steps |
| `client.SequencesAPI.UpdateSequence(ctx)` | Update sequence |
| `client.SequencesAPI.DeleteSequence(ctx)` | Delete sequence |
| `client.SequencesAPI.ActivateSequence(ctx)` | Activate sequence |
| `client.SequencesAPI.EnrollContacts(ctx)` | Enroll contacts in a sequence |
| `client.SequencesAPI.PauseSequence(ctx)` | Pause sequence |
| `client.SequencesAPI.UnenrollContact(ctx)` | Unenroll contact |

### Slack
| Method | Description |
|--------|-------------|
| `client.SlackAPI.ListSlackMembers(ctx)` | List Slack workspace members |

### Tracking Tags
| Method | Description |
|--------|-------------|
| `client.TrackingTagsAPI.ListTrackingTagEvents(ctx)` | List conversion events |
| `client.TrackingTagsAPI.ListTrackingTagPartners(ctx)` | List partner businesses of a tag |
| `client.TrackingTagsAPI.ListTrackingTagSharedAccounts(ctx)` | List accounts it is shared with |
| `client.TrackingTagsAPI.ListTrackingTagUsers(ctx)` | List tag users |
| `client.TrackingTagsAPI.ListTrackingTags(ctx)` | List tracking tags |
| `client.TrackingTagsAPI.CreateTrackingTag(ctx)` | Create a tracking tag |
| `client.TrackingTagsAPI.CreateTrackingTagEvent(ctx)` | Create a conversion event |
| `client.TrackingTagsAPI.GetAdTrackingTags(ctx)` | Get ad tracking tags |
| `client.TrackingTagsAPI.GetTrackingTag(ctx)` | Get a tracking tag |
| `client.TrackingTagsAPI.GetTrackingTagDiagnostics(ctx)` | Get tag diagnostics |
| `client.TrackingTagsAPI.GetTrackingTagStats(ctx)` | Get aggregated event stats |
| `client.TrackingTagsAPI.GetTrackingTagStoreInstall(ctx)` | Get store install status |
| `client.TrackingTagsAPI.UpdateAdTrackingTags(ctx)` | Set ad tracking tags |
| `client.TrackingTagsAPI.UpdateTrackingTag(ctx)` | Update a tracking tag |
| `client.TrackingTagsAPI.UpdateTrackingTagEvent(ctx)` | Update a conversion event |
| `client.TrackingTagsAPI.DeleteTrackingTagEvent(ctx)` | Delete a conversion event |
| `client.TrackingTagsAPI.AddTrackingTagSharedAccount(ctx)` | Share with an ad account |
| `client.TrackingTagsAPI.AssignTrackingTagUser(ctx)` | Assign a user to a tag |
| `client.TrackingTagsAPI.InstallTrackingTagOnStore(ctx)` | Install on a Shopify store or WordPress site |
| `client.TrackingTagsAPI.RemoveTrackingTagFromStore(ctx)` | Remove from a Shopify store or WordPress site |
| `client.TrackingTagsAPI.RemoveTrackingTagSharedAccount(ctx)` | Stop sharing with an account |
| `client.TrackingTagsAPI.RemoveTrackingTagUser(ctx)` | Remove a user from a tag |

### Twitter Engagement
| Method | Description |
|--------|-------------|
| `client.TwitterEngagementAPI.GetTweet(ctx)` | Look up a tweet |
| `client.TwitterEngagementAPI.BookmarkPost(ctx)` | Bookmark a tweet |
| `client.TwitterEngagementAPI.FollowUser(ctx)` | Follow a user |
| `client.TwitterEngagementAPI.RemoveBookmark(ctx)` | Remove bookmark |
| `client.TwitterEngagementAPI.RetweetPost(ctx)` | Retweet a post |
| `client.TwitterEngagementAPI.SearchTweets(ctx)` | Search recent tweets |
| `client.TwitterEngagementAPI.UndoRetweet(ctx)` | Undo retweet |
| `client.TwitterEngagementAPI.UnfollowUser(ctx)` | Unfollow a user |

### Validate
| Method | Description |
|--------|-------------|
| `client.ValidateAPI.ValidateMedia(ctx)` | Validate media URL |
| `client.ValidateAPI.ValidatePost(ctx)` | Validate post content |
| `client.ValidateAPI.ValidatePostLength(ctx)` | Validate character count |
| `client.ValidateAPI.ValidateSubreddit(ctx)` | Check subreddit existence |

### Verify
| Method | Description |
|--------|-------------|
| `client.VerifyAPI.CreateVerification(ctx)` | Send a verification code |
| `client.VerifyAPI.GetVerification(ctx)` | Get a verification |
| `client.VerifyAPI.CheckVerification(ctx)` | Check a verification code |

### Voice
| Method | Description |
|--------|-------------|
| `client.VoiceAPI.ListSipTrunks(ctx)` | List SIP trunks |
| `client.VoiceAPI.ListVoiceCalls(ctx)` | List phone calls |
| `client.VoiceAPI.CreateSipTrunk(ctx)` | Create a SIP trunk |
| `client.VoiceAPI.CreateVoiceCall(ctx)` | Place an outbound phone call |
| `client.VoiceAPI.CreateVoiceWebSession(ctx)` | Mint a browser softphone session |
| `client.VoiceAPI.GetSipTrunk(ctx)` | Get a SIP trunk |
| `client.VoiceAPI.GetVoiceCall(ctx)` | Get a phone call |
| `client.VoiceAPI.GetVoiceCallEstimate(ctx)` | Estimate call cost |
| `client.VoiceAPI.GetVoiceCallRecording(ctx)` | Get a call recording |
| `client.VoiceAPI.DeleteSipTrunk(ctx)` | Delete a SIP trunk |
| `client.VoiceAPI.AttachNumberToSipTrunk(ctx)` | Attach a number to a SIP trunk |
| `client.VoiceAPI.DetachNumberFromSipTrunk(ctx)` | Detach a number from its SIP trunk |
| `client.VoiceAPI.DialVoiceWebCall(ctx)` | Dial from the browser softphone |
| `client.VoiceAPI.DisableVoiceOnNumber(ctx)` | Disable phone calling on a number |
| `client.VoiceAPI.EnableVoiceOnNumber(ctx)` | Enable phone calling on a number |
| `client.VoiceAPI.EndVoiceCall(ctx)` | Hang up a live call |
| `client.VoiceAPI.RotateSipTrunkCredentials(ctx)` | Rotate a SIP trunk's password |
| `client.VoiceAPI.TransferVoiceCall(ctx)` | Blind-transfer a live call |

### WhatsApp
| Method | Description |
|--------|-------------|
| `client.WhatsAppAPI.ListWhatsAppAccountEvents(ctx)` | List account notifications |
| `client.WhatsAppAPI.ListWhatsAppCatalogs(ctx)` | List the catalogs linked to a WhatsApp number |
| `client.WhatsAppAPI.ListWhatsAppConversions(ctx)` | List conversion events |
| `client.WhatsAppAPI.ListWhatsAppGroupChats(ctx)` | List active groups |
| `client.WhatsAppAPI.ListWhatsAppGroupJoinRequests(ctx)` | List join requests |
| `client.WhatsAppAPI.CreateWhatsAppDataset(ctx)` | Provision CTWA dataset |
| `client.WhatsAppAPI.CreateWhatsAppGroupChat(ctx)` | Create group |
| `client.WhatsAppAPI.CreateWhatsAppGroupInviteLink(ctx)` | Create invite link |
| `client.WhatsAppAPI.CreateWhatsAppTemplate(ctx)` | Create template |
| `client.WhatsAppAPI.GetWhatsAppBlockStatus(ctx)` | Check if a user is blocked |
| `client.WhatsAppAPI.GetWhatsAppBlockedUsers(ctx)` | List blocked users |
| `client.WhatsAppAPI.GetWhatsAppBusinessProfile(ctx)` | Get business profile |
| `client.WhatsAppAPI.GetWhatsAppCommerceSettings(ctx)` | Get a number's commerce settings |
| `client.WhatsAppAPI.GetWhatsAppDataset(ctx)` | Get CTWA conversions dataset |
| `client.WhatsAppAPI.GetWhatsAppDisplayName(ctx)` | Get display name status |
| `client.WhatsAppAPI.GetWhatsAppGroupChat(ctx)` | Get group info |
| `client.WhatsAppAPI.GetWhatsAppMedia(ctx)` | Download WhatsApp media |
| `client.WhatsAppAPI.GetWhatsAppTemplate(ctx)` | Get template |
| `client.WhatsAppAPI.GetWhatsAppTemplateById(ctx)` | Get template by id |
| `client.WhatsAppAPI.GetWhatsAppTemplates(ctx)` | List templates |
| `client.WhatsAppAPI.GetWhatsappBusinessUsername(ctx)` | Get business username |
| `client.WhatsAppAPI.GetWhatsappBusinessUsernameSuggestions(ctx)` | Get username suggestions |
| `client.WhatsAppAPI.UpdateWhatsAppBusinessProfile(ctx)` | Update business profile |
| `client.WhatsAppAPI.UpdateWhatsAppCommerceSettings(ctx)` | Update a number's commerce settings |
| `client.WhatsAppAPI.UpdateWhatsAppDisplayName(ctx)` | Request display name change |
| `client.WhatsAppAPI.UpdateWhatsAppGroupChat(ctx)` | Update group settings |
| `client.WhatsAppAPI.UpdateWhatsAppTemplate(ctx)` | Update template |
| `client.WhatsAppAPI.UpdateWhatsAppTemplateById(ctx)` | Update template by id |
| `client.WhatsAppAPI.DeleteWhatsAppGroupChat(ctx)` | Delete group |
| `client.WhatsAppAPI.DeleteWhatsAppTemplate(ctx)` | Delete template |
| `client.WhatsAppAPI.DeleteWhatsAppTemplateById(ctx)` | Delete template by id |
| `client.WhatsAppAPI.DeleteWhatsappBusinessUsername(ctx)` | Delete business username |
| `client.WhatsAppAPI.AddWhatsAppGroupParticipants(ctx)` | Add participants |
| `client.WhatsAppAPI.ApproveWhatsAppGroupJoinRequests(ctx)` | Approve join requests |
| `client.WhatsAppAPI.BlockWhatsAppUsers(ctx)` | Block users |
| `client.WhatsAppAPI.LinkWhatsAppCatalog(ctx)` | Link a catalog to a WhatsApp number |
| `client.WhatsAppAPI.RegisterWhatsAppNumber(ctx)` | Register a connected WhatsApp number on the Cloud API |
| `client.WhatsAppAPI.RejectWhatsAppGroupJoinRequests(ctx)` | Reject join requests |
| `client.WhatsAppAPI.RemoveWhatsAppGroupParticipants(ctx)` | Remove participants |
| `client.WhatsAppAPI.RequestWhatsAppVerificationCode(ctx)` | Request a Meta re-verification code for a BYO WhatsApp number |
| `client.WhatsAppAPI.SendWhatsAppConversion(ctx)` | Send WhatsApp conversion event |
| `client.WhatsAppAPI.SetWhatsappBusinessUsername(ctx)` | Set business username |
| `client.WhatsAppAPI.UnblockWhatsAppUsers(ctx)` | Unblock users |
| `client.WhatsAppAPI.UnlinkWhatsAppCatalog(ctx)` | Unlink a catalog from a WhatsApp number |
| `client.WhatsAppAPI.UploadWhatsAppProfilePhoto(ctx)` | Upload profile picture |
| `client.WhatsAppAPI.VerifyWhatsAppNumber(ctx)` | Verify the Meta re-verification code for a BYO WhatsApp number |

### WhatsApp Calling
| Method | Description |
|--------|-------------|
| `client.WhatsAppCallingAPI.ListWhatsAppCalls(ctx)` | List call history for an account |
| `client.WhatsAppCallingAPI.GetWhatsAppCall(ctx)` | Get a single call |
| `client.WhatsAppCallingAPI.GetWhatsAppCallEstimate(ctx)` | Estimate per-minute cost |
| `client.WhatsAppCallingAPI.GetWhatsAppCallPermissions(ctx)` | Check call permission |
| `client.WhatsAppCallingAPI.GetWhatsAppCallRecording(ctx)` | Get a call recording |
| `client.WhatsAppCallingAPI.GetWhatsAppCalling(ctx)` | Get calling config for a number |
| `client.WhatsAppCallingAPI.GetWhatsAppCallingConfig(ctx)` | Get calling config for an account |
| `client.WhatsAppCallingAPI.UpdateWhatsAppCalling(ctx)` | Update calling config |
| `client.WhatsAppCallingAPI.UpdateWhatsAppCallingLegacy(ctx)` | Update calling config |
| `client.WhatsAppCallingAPI.DisableWhatsAppCalling(ctx)` | Disable calling on a number |
| `client.WhatsAppCallingAPI.DisableWhatsAppCallingLegacy(ctx)` | Disable calling on a number |
| `client.WhatsAppCallingAPI.EnableWhatsAppCalling(ctx)` | Enable calling on a number |
| `client.WhatsAppCallingAPI.EnableWhatsAppCallingLegacy(ctx)` | Enable calling on a number |
| `client.WhatsAppCallingAPI.InitiateWhatsAppCall(ctx)` | Initiate outbound call |
| `client.WhatsAppCallingAPI.StartWhatsAppCallerIdVerification(ctx)` | Start caller-ID verification for a customer-brought number |
| `client.WhatsAppCallingAPI.VerifyWhatsAppCallerId(ctx)` | Confirm the caller-ID verification code |

### WhatsApp Flows
| Method | Description |
|--------|-------------|
| `client.WhatsAppFlowsAPI.ListWhatsAppFlowResponses(ctx)` | List flow responses |
| `client.WhatsAppFlowsAPI.ListWhatsAppFlowVersions(ctx)` | List flow versions |
| `client.WhatsAppFlowsAPI.ListWhatsAppFlows(ctx)` | List flows |
| `client.WhatsAppFlowsAPI.CreateWhatsAppFlow(ctx)` | Create flow |
| `client.WhatsAppFlowsAPI.GetWhatsAppFlow(ctx)` | Get flow |
| `client.WhatsAppFlowsAPI.GetWhatsAppFlowJson(ctx)` | Get flow JSON asset |
| `client.WhatsAppFlowsAPI.GetWhatsAppFlowPreview(ctx)` | Get flow preview URL |
| `client.WhatsAppFlowsAPI.GetWhatsAppFlowsEncryptionKey(ctx)` | Get Flows encryption key status |
| `client.WhatsAppFlowsAPI.UpdateWhatsAppFlow(ctx)` | Update flow |
| `client.WhatsAppFlowsAPI.DeleteWhatsAppFlow(ctx)` | Delete flow |
| `client.WhatsAppFlowsAPI.DeprecateWhatsAppFlow(ctx)` | Deprecate flow |
| `client.WhatsAppFlowsAPI.PublishWhatsAppFlow(ctx)` | Publish flow |
| `client.WhatsAppFlowsAPI.SendWhatsAppFlowMessage(ctx)` | Send flow message |
| `client.WhatsAppFlowsAPI.SetWhatsAppFlowsEncryptionKey(ctx)` | Register a Flows encryption key |
| `client.WhatsAppFlowsAPI.UploadWhatsAppFlowJson(ctx)` | Upload flow JSON |

### WhatsApp Phone Numbers
| Method | Description |
|--------|-------------|
| `client.WhatsAppPhoneNumbersAPI.ListWhatsAppNumberCountries(ctx)` | List offerable number countries |
| `client.WhatsAppPhoneNumbersAPI.CreateWhatsAppNumberKycLink(ctx)` | Create a hosted KYC link |
| `client.WhatsAppPhoneNumbersAPI.GetWhatsAppNumberInfo(ctx)` | Get number status |
| `client.WhatsAppPhoneNumbersAPI.GetWhatsAppNumberKycForm(ctx)` | Get KYC form spec |
| `client.WhatsAppPhoneNumbersAPI.GetWhatsAppNumberRemediation(ctx)` | Get declined requirements |
| `client.WhatsAppPhoneNumbersAPI.GetWhatsAppPhoneNumber(ctx)` | Get phone number |
| `client.WhatsAppPhoneNumbersAPI.GetWhatsAppPhoneNumbers(ctx)` | List phone numbers |
| `client.WhatsAppPhoneNumbersAPI.GetWhatsAppPricingAnalytics(ctx)` | Get pricing analytics |
| `client.WhatsAppPhoneNumbersAPI.CheckWhatsAppNumberAvailability(ctx)` | Check country availability |
| `client.WhatsAppPhoneNumbersAPI.MoveWhatsAppNumberToProfile(ctx)` | Move a number to another profile |
| `client.WhatsAppPhoneNumbersAPI.PurchaseWhatsAppPhoneNumber(ctx)` | Purchase phone number |
| `client.WhatsAppPhoneNumbersAPI.ReleaseWhatsAppPhoneNumber(ctx)` | Release phone number |
| `client.WhatsAppPhoneNumbersAPI.RemediateWhatsAppNumber(ctx)` | Resubmit a declined number |
| `client.WhatsAppPhoneNumbersAPI.SearchAvailableWhatsAppNumbers(ctx)` | Search available numbers |
| `client.WhatsAppPhoneNumbersAPI.SubmitWhatsAppNumberKyc(ctx)` | Submit KYC |
| `client.WhatsAppPhoneNumbersAPI.UploadWhatsAppNumberKycDocument(ctx)` | Upload a KYC document |
| `client.WhatsAppPhoneNumbersAPI.ValidateWhatsAppNumberKycAddress(ctx)` | Pre-validate KYC address |

### WhatsApp Sandbox
| Method | Description |
|--------|-------------|
| `client.WhatsAppSandboxAPI.ListWhatsAppSandboxSessions(ctx)` | List your sandbox sessions |
| `client.WhatsAppSandboxAPI.CreateWhatsAppSandboxSession(ctx)` | Start a sandbox activation |
| `client.WhatsAppSandboxAPI.DeleteWhatsAppSandboxSession(ctx)` | Revoke a sandbox session |

### WhatsApp Templates
| Method | Description |
|--------|-------------|
| `client.WhatsAppTemplatesAPI.GetWhatsAppLibraryTemplate(ctx)` | Look up a library template |

### Workflows
| Method | Description |
|--------|-------------|
| `client.WorkflowsAPI.ListWorkflowExecutionEvents(ctx)` | Get an execution's timeline |
| `client.WorkflowsAPI.ListWorkflowExecutions(ctx)` | List workflow runs |
| `client.WorkflowsAPI.ListWorkflowVersions(ctx)` | List a workflow's version history |
| `client.WorkflowsAPI.ListWorkflows(ctx)` | List workflows |
| `client.WorkflowsAPI.CreateWorkflow(ctx)` | Create workflow |
| `client.WorkflowsAPI.GetWorkflow(ctx)` | Get workflow with graph |
| `client.WorkflowsAPI.GetWorkflowVersion(ctx)` | Get a specific workflow version |
| `client.WorkflowsAPI.UpdateWorkflow(ctx)` | Update workflow |
| `client.WorkflowsAPI.DeleteWorkflow(ctx)` | Delete workflow |
| `client.WorkflowsAPI.ActivateWorkflow(ctx)` | Activate workflow |
| `client.WorkflowsAPI.DuplicateWorkflow(ctx)` | Duplicate a workflow |
| `client.WorkflowsAPI.PauseWorkflow(ctx)` | Pause workflow |
| `client.WorkflowsAPI.RestoreWorkflowVersion(ctx)` | Restore a workflow version |
| `client.WorkflowsAPI.TriggerWorkflow(ctx)` | Manually start a workflow run |

### iMessage
| Method | Description |
|--------|-------------|
| `client.IMessageAPI.ListImessageAudience(ctx)` | List iMessage audience |
| `client.IMessageAPI.ListImessageAvailableNumbers(ctx)` | List instantly available iMessage numbers |
| `client.IMessageAPI.ListImessageSandboxContacts(ctx)` | List iMessage sandbox contacts |
| `client.IMessageAPI.ListImessageSenderOrders(ctx)` | List iMessage sender orders |
| `client.IMessageAPI.ListImessageSenders(ctx)` | List iMessage senders |
| `client.IMessageAPI.CreateImessageGroup(ctx)` | Start an iMessage group chat |
| `client.IMessageAPI.CreateImessageOptInLink(ctx)` | Create a tracked iMessage opt-in link |
| `client.IMessageAPI.GetImessageGroup(ctx)` | Get an iMessage group |
| `client.IMessageAPI.GetImessageSender(ctx)` | Get iMessage sender status |
| `client.IMessageAPI.UpdateImessageGroup(ctx)` | Rename an iMessage group or change its photo |
| `client.IMessageAPI.UpdateImessageSender(ctx)` | Update an iMessage sender |
| `client.IMessageAPI.AddImessageGroupParticipant(ctx)` | Add a participant to an iMessage group |
| `client.IMessageAPI.AddImessageSandboxContact(ctx)` | Add an iMessage sandbox contact |
| `client.IMessageAPI.CancelImessageSender(ctx)` | Cancel an iMessage sender |
| `client.IMessageAPI.OrderImessageSender(ctx)` | Order a new iMessage sender |
| `client.IMessageAPI.RegisterImessageSender(ctx)` | Register an iMessage sender |
| `client.IMessageAPI.RemoveImessageGroupParticipant(ctx)` | Remove a participant from an iMessage group |
| `client.IMessageAPI.RemoveImessageSandboxContact(ctx)` | Remove an iMessage sandbox contact |
| `client.IMessageAPI.ReserveImessageAvailableNumber(ctx)` | Reserve an available iMessage number |
| `client.IMessageAPI.SetImessageSubscription(ctx)` | Subscribe or opt out an iMessage contact |

### Invites
| Method | Description |
|--------|-------------|
| `client.InvitesAPI.CreateInviteToken(ctx)` | Create invite token |

## Documentation

- [API Reference](https://docs.zernio.com/api-reference)
- [Getting Started Guide](https://docs.zernio.com/quickstart)

## License

MIT
