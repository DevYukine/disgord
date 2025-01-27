// Code generated by generate/interfaces; DO NOT EDIT.

package disgord

import (
	"context"
	"github.com/andersfylling/disgord/internal/gateway"
	"net/url"
)

type applicationCommandQueryBuilderNop struct {
	Ctx       context.Context
	Flags     Flag
	ChannelID Snowflake
	GuildID   Snowflake
	UserID    Snowflake
}

var _ ApplicationCommandQueryBuilder = &applicationCommandQueryBuilderNop{}

func (a applicationCommandQueryBuilderNop) WithContext(ctx context.Context) ApplicationCommandQueryBuilder {
	a.Ctx = ctx
	return &a
}

func (a *applicationCommandQueryBuilderNop) Global() ApplicationCommandFunctions {
	return nil
}

func (a *applicationCommandQueryBuilderNop) Guild(_ Snowflake) ApplicationCommandFunctions {
	return nil
}

type channelQueryBuilderNop struct {
	Ctx       context.Context
	Flags     Flag
	ChannelID Snowflake
	GuildID   Snowflake
	UserID    Snowflake
}

var _ ChannelQueryBuilder = &channelQueryBuilderNop{}

func (c channelQueryBuilderNop) WithContext(ctx context.Context) ChannelQueryBuilder {
	c.Ctx = ctx
	return &c
}

func (c channelQueryBuilderNop) WithFlags(flags ...Flag) ChannelQueryBuilder {
	c.Flags = mergeFlags(flags)
	return &c
}

func (c *channelQueryBuilderNop) AddDMParticipant(_ *GroupDMParticipant) error {
	return nil
}

func (c *channelQueryBuilderNop) AddThreadMember(_ Snowflake) error {
	return nil
}

func (c *channelQueryBuilderNop) CreateInvite(_ *CreateInvite) (*Invite, error) {
	return nil, nil
}

func (c *channelQueryBuilderNop) CreateMessage(_ *CreateMessage) (*Message, error) {
	return nil, nil
}

func (c *channelQueryBuilderNop) CreateThread(_ *CreateThreadWithoutMessage) (*Channel, error) {
	return nil, nil
}

func (c *channelQueryBuilderNop) CreateWebhook(_ *CreateWebhook) (*Webhook, error) {
	return nil, nil
}

func (c *channelQueryBuilderNop) Delete() (*Channel, error) {
	return nil, nil
}

func (c *channelQueryBuilderNop) DeleteMessages(_ *DeleteMessages) error {
	return nil
}

func (c *channelQueryBuilderNop) DeletePermission(_ Snowflake) error {
	return nil
}

func (c *channelQueryBuilderNop) Get() (*Channel, error) {
	return nil, nil
}

func (c *channelQueryBuilderNop) GetInvites() ([]*Invite, error) {
	return nil, nil
}

func (c *channelQueryBuilderNop) GetJoinedPrivateArchivedThreads(_ *GetArchivedThreads) (*ArchivedThreads, error) {
	return nil, nil
}

func (c *channelQueryBuilderNop) GetMessages(_ *GetMessages) ([]*Message, error) {
	return nil, nil
}

func (c *channelQueryBuilderNop) GetPinnedMessages() ([]*Message, error) {
	return nil, nil
}

func (c *channelQueryBuilderNop) GetPrivateArchivedThreads(_ *GetArchivedThreads) (*ArchivedThreads, error) {
	return nil, nil
}

func (c *channelQueryBuilderNop) GetPublicArchivedThreads(_ *GetArchivedThreads) (*ArchivedThreads, error) {
	return nil, nil
}

func (c *channelQueryBuilderNop) GetThreadMember(_ Snowflake) (*ThreadMember, error) {
	return nil, nil
}

func (c *channelQueryBuilderNop) GetThreadMembers() ([]*ThreadMember, error) {
	return nil, nil
}

func (c *channelQueryBuilderNop) GetWebhooks() ([]*Webhook, error) {
	return nil, nil
}

func (c *channelQueryBuilderNop) JoinThread() error {
	return nil
}

func (c *channelQueryBuilderNop) KickParticipant(_ Snowflake) error {
	return nil
}

func (c *channelQueryBuilderNop) LeaveThread() error {
	return nil
}

func (c *channelQueryBuilderNop) Message(_ Snowflake) MessageQueryBuilder {
	return nil
}

func (c *channelQueryBuilderNop) RemoveThreadMember(_ Snowflake) error {
	return nil
}

func (c *channelQueryBuilderNop) TriggerTypingIndicator() error {
	return nil
}

func (c *channelQueryBuilderNop) Update(_ *UpdateChannel) (*Channel, error) {
	return nil, nil
}

func (c *channelQueryBuilderNop) UpdateBuilder() UpdateChannelBuilder {
	return nil
}

func (c *channelQueryBuilderNop) UpdatePermissions(_ Snowflake, _ *UpdateChannelPermissions) error {
	return nil
}

type clientQueryBuilderNop struct {
	Ctx       context.Context
	Flags     Flag
	ChannelID Snowflake
	GuildID   Snowflake
	UserID    Snowflake
}

var _ ClientQueryBuilder = &clientQueryBuilderNop{}

func (c clientQueryBuilderNop) WithContext(ctx context.Context) ClientQueryBuilderExecutables {
	c.Ctx = ctx
	return &c
}

func (c clientQueryBuilderNop) WithFlags(flags ...Flag) ClientQueryBuilderExecutables {
	c.Flags = mergeFlags(flags)
	return &c
}

func (c *clientQueryBuilderNop) ApplicationCommand(_ Snowflake) ApplicationCommandQueryBuilder {
	return nil
}

func (c *clientQueryBuilderNop) BotAuthorizeURL(_ PermissionBit, _ []string) (*url.URL, error) {
	return nil, nil
}

func (c *clientQueryBuilderNop) Channel(_ Snowflake) ChannelQueryBuilder {
	return nil
}

func (c *clientQueryBuilderNop) CreateGuild(_ string, _ *CreateGuild) (*Guild, error) {
	return nil, nil
}

func (c *clientQueryBuilderNop) CurrentUser() CurrentUserQueryBuilder {
	return nil
}

func (c *clientQueryBuilderNop) Gateway() GatewayQueryBuilder {
	return nil
}

func (c *clientQueryBuilderNop) GetVoiceRegions() ([]*VoiceRegion, error) {
	return nil, nil
}

func (c *clientQueryBuilderNop) Guild(_ Snowflake) GuildQueryBuilder {
	return nil
}

func (c *clientQueryBuilderNop) Invite(_ string) InviteQueryBuilder {
	return nil
}

func (c *clientQueryBuilderNop) SendMsg(_ Snowflake, _ ...interface{}) (*Message, error) {
	return nil, nil
}

func (c *clientQueryBuilderNop) User(_ Snowflake) UserQueryBuilder {
	return nil
}

type currentUserQueryBuilderNop struct {
	Ctx       context.Context
	Flags     Flag
	ChannelID Snowflake
	GuildID   Snowflake
	UserID    Snowflake
}

var _ CurrentUserQueryBuilder = &currentUserQueryBuilderNop{}

func (c currentUserQueryBuilderNop) WithContext(ctx context.Context) CurrentUserQueryBuilder {
	c.Ctx = ctx
	return &c
}

func (c currentUserQueryBuilderNop) WithFlags(flags ...Flag) CurrentUserQueryBuilder {
	c.Flags = mergeFlags(flags)
	return &c
}

func (c *currentUserQueryBuilderNop) CreateGroupDM(_ *CreateGroupDM) (*Channel, error) {
	return nil, nil
}

func (c *currentUserQueryBuilderNop) Get() (*User, error) {
	return nil, nil
}

func (c *currentUserQueryBuilderNop) GetConnections() ([]*UserConnection, error) {
	return nil, nil
}

func (c *currentUserQueryBuilderNop) GetGuilds(_ *GetCurrentUserGuilds) ([]*Guild, error) {
	return nil, nil
}

func (c *currentUserQueryBuilderNop) GetUserConnections() ([]*UserConnection, error) {
	return nil, nil
}

func (c *currentUserQueryBuilderNop) LeaveGuild(_ Snowflake) error {
	return nil
}

func (c *currentUserQueryBuilderNop) Update(_ *UpdateUser) (*User, error) {
	return nil, nil
}

func (c *currentUserQueryBuilderNop) UpdateBuilder() UpdateCurrentUserBuilder {
	return nil
}

type gatewayQueryBuilderNop struct {
	Ctx       context.Context
	Flags     Flag
	ChannelID Snowflake
	GuildID   Snowflake
	UserID    Snowflake
}

var _ GatewayQueryBuilder = &gatewayQueryBuilderNop{}

func (g gatewayQueryBuilderNop) WithContext(ctx context.Context) GatewayQueryBuilder {
	g.Ctx = ctx
	return &g
}

func (g *gatewayQueryBuilderNop) BotGuildsReady(_ func()) {
	return
}

func (g *gatewayQueryBuilderNop) BotReady(_ func()) {
	return
}

func (g *gatewayQueryBuilderNop) ChannelCreate(_ func(Session, *ChannelCreate), _ ...func(Session, *ChannelCreate)) {
	return
}

func (g *gatewayQueryBuilderNop) ChannelCreateChan(_ chan *ChannelCreate, _ ...chan *ChannelCreate) {
	return
}

func (g *gatewayQueryBuilderNop) ChannelDelete(_ func(Session, *ChannelDelete), _ ...func(Session, *ChannelDelete)) {
	return
}

func (g *gatewayQueryBuilderNop) ChannelDeleteChan(_ chan *ChannelDelete, _ ...chan *ChannelDelete) {
	return
}

func (g *gatewayQueryBuilderNop) ChannelPinsUpdate(_ func(Session, *ChannelPinsUpdate), _ ...func(Session, *ChannelPinsUpdate)) {
	return
}

func (g *gatewayQueryBuilderNop) ChannelPinsUpdateChan(_ chan *ChannelPinsUpdate, _ ...chan *ChannelPinsUpdate) {
	return
}

func (g *gatewayQueryBuilderNop) ChannelUpdate(_ func(Session, *ChannelUpdate), _ ...func(Session, *ChannelUpdate)) {
	return
}

func (g *gatewayQueryBuilderNop) ChannelUpdateChan(_ chan *ChannelUpdate, _ ...chan *ChannelUpdate) {
	return
}

func (g *gatewayQueryBuilderNop) Connect() error {
	return nil
}

func (g *gatewayQueryBuilderNop) Disconnect() error {
	return nil
}

func (g *gatewayQueryBuilderNop) DisconnectOnInterrupt() error {
	return nil
}

func (g *gatewayQueryBuilderNop) Dispatch(_ GatewayCmdName, _ gateway.CmdPayload) ([]Snowflake, error) {
	return nil, nil
}

func (g *gatewayQueryBuilderNop) Get() (*gateway.Gateway, error) {
	return nil, nil
}

func (g *gatewayQueryBuilderNop) GetBot() (*gateway.GatewayBot, error) {
	return nil, nil
}

func (g *gatewayQueryBuilderNop) GuildBanAdd(_ func(Session, *GuildBanAdd), _ ...func(Session, *GuildBanAdd)) {
	return
}

func (g *gatewayQueryBuilderNop) GuildBanAddChan(_ chan *GuildBanAdd, _ ...chan *GuildBanAdd) {
	return
}

func (g *gatewayQueryBuilderNop) GuildBanRemove(_ func(Session, *GuildBanRemove), _ ...func(Session, *GuildBanRemove)) {
	return
}

func (g *gatewayQueryBuilderNop) GuildBanRemoveChan(_ chan *GuildBanRemove, _ ...chan *GuildBanRemove) {
	return
}

func (g *gatewayQueryBuilderNop) GuildCreate(_ func(Session, *GuildCreate), _ ...func(Session, *GuildCreate)) {
	return
}

func (g *gatewayQueryBuilderNop) GuildCreateChan(_ chan *GuildCreate, _ ...chan *GuildCreate) {
	return
}

func (g *gatewayQueryBuilderNop) GuildDelete(_ func(Session, *GuildDelete), _ ...func(Session, *GuildDelete)) {
	return
}

func (g *gatewayQueryBuilderNop) GuildDeleteChan(_ chan *GuildDelete, _ ...chan *GuildDelete) {
	return
}

func (g *gatewayQueryBuilderNop) GuildEmojisUpdate(_ func(Session, *GuildEmojisUpdate), _ ...func(Session, *GuildEmojisUpdate)) {
	return
}

func (g *gatewayQueryBuilderNop) GuildEmojisUpdateChan(_ chan *GuildEmojisUpdate, _ ...chan *GuildEmojisUpdate) {
	return
}

func (g *gatewayQueryBuilderNop) GuildIntegrationsUpdate(_ func(Session, *GuildIntegrationsUpdate), _ ...func(Session, *GuildIntegrationsUpdate)) {
	return
}

func (g *gatewayQueryBuilderNop) GuildIntegrationsUpdateChan(_ chan *GuildIntegrationsUpdate, _ ...chan *GuildIntegrationsUpdate) {
	return
}

func (g *gatewayQueryBuilderNop) GuildMemberAdd(_ func(Session, *GuildMemberAdd), _ ...func(Session, *GuildMemberAdd)) {
	return
}

func (g *gatewayQueryBuilderNop) GuildMemberAddChan(_ chan *GuildMemberAdd, _ ...chan *GuildMemberAdd) {
	return
}

func (g *gatewayQueryBuilderNop) GuildMemberRemove(_ func(Session, *GuildMemberRemove), _ ...func(Session, *GuildMemberRemove)) {
	return
}

func (g *gatewayQueryBuilderNop) GuildMemberRemoveChan(_ chan *GuildMemberRemove, _ ...chan *GuildMemberRemove) {
	return
}

func (g *gatewayQueryBuilderNop) GuildMembersChunk(_ func(Session, *GuildMembersChunk), _ ...func(Session, *GuildMembersChunk)) {
	return
}

func (g *gatewayQueryBuilderNop) GuildMembersChunkChan(_ chan *GuildMembersChunk, _ ...chan *GuildMembersChunk) {
	return
}

func (g *gatewayQueryBuilderNop) GuildMemberUpdate(_ func(Session, *GuildMemberUpdate), _ ...func(Session, *GuildMemberUpdate)) {
	return
}

func (g *gatewayQueryBuilderNop) GuildMemberUpdateChan(_ chan *GuildMemberUpdate, _ ...chan *GuildMemberUpdate) {
	return
}

func (g *gatewayQueryBuilderNop) GuildRoleCreate(_ func(Session, *GuildRoleCreate), _ ...func(Session, *GuildRoleCreate)) {
	return
}

func (g *gatewayQueryBuilderNop) GuildRoleCreateChan(_ chan *GuildRoleCreate, _ ...chan *GuildRoleCreate) {
	return
}

func (g *gatewayQueryBuilderNop) GuildRoleDelete(_ func(Session, *GuildRoleDelete), _ ...func(Session, *GuildRoleDelete)) {
	return
}

func (g *gatewayQueryBuilderNop) GuildRoleDeleteChan(_ chan *GuildRoleDelete, _ ...chan *GuildRoleDelete) {
	return
}

func (g *gatewayQueryBuilderNop) GuildRoleUpdate(_ func(Session, *GuildRoleUpdate), _ ...func(Session, *GuildRoleUpdate)) {
	return
}

func (g *gatewayQueryBuilderNop) GuildRoleUpdateChan(_ chan *GuildRoleUpdate, _ ...chan *GuildRoleUpdate) {
	return
}

func (g *gatewayQueryBuilderNop) GuildScheduledEventCreate(_ func(Session, *GuildScheduledEventCreate), _ ...func(Session, *GuildScheduledEventCreate)) {
	return
}

func (g *gatewayQueryBuilderNop) GuildScheduledEventCreateChan(_ chan *GuildScheduledEventCreate, _ ...chan *GuildScheduledEventCreate) {
	return
}

func (g *gatewayQueryBuilderNop) GuildScheduledEventDelete(_ func(Session, *GuildScheduledEventDelete), _ ...func(Session, *GuildScheduledEventDelete)) {
	return
}

func (g *gatewayQueryBuilderNop) GuildScheduledEventDeleteChan(_ chan *GuildScheduledEventDelete, _ ...chan *GuildScheduledEventDelete) {
	return
}

func (g *gatewayQueryBuilderNop) GuildScheduledEventUpdate(_ func(Session, *GuildScheduledEventUpdate), _ ...func(Session, *GuildScheduledEventUpdate)) {
	return
}

func (g *gatewayQueryBuilderNop) GuildScheduledEventUpdateChan(_ chan *GuildScheduledEventUpdate, _ ...chan *GuildScheduledEventUpdate) {
	return
}

func (g *gatewayQueryBuilderNop) GuildScheduledEventUserAdd(_ func(Session, *GuildScheduledEventUserAdd), _ ...func(Session, *GuildScheduledEventUserAdd)) {
	return
}

func (g *gatewayQueryBuilderNop) GuildScheduledEventUserAddChan(_ chan *GuildScheduledEventUserAdd, _ ...chan *GuildScheduledEventUserAdd) {
	return
}

func (g *gatewayQueryBuilderNop) GuildScheduledEventUserRemove(_ func(Session, *GuildScheduledEventUserRemove), _ ...func(Session, *GuildScheduledEventUserRemove)) {
	return
}

func (g *gatewayQueryBuilderNop) GuildScheduledEventUserRemoveChan(_ chan *GuildScheduledEventUserRemove, _ ...chan *GuildScheduledEventUserRemove) {
	return
}

func (g *gatewayQueryBuilderNop) GuildStickersUpdate(_ func(Session, *GuildStickersUpdate), _ ...func(Session, *GuildStickersUpdate)) {
	return
}

func (g *gatewayQueryBuilderNop) GuildStickersUpdateChan(_ chan *GuildStickersUpdate, _ ...chan *GuildStickersUpdate) {
	return
}

func (g *gatewayQueryBuilderNop) GuildUpdate(_ func(Session, *GuildUpdate), _ ...func(Session, *GuildUpdate)) {
	return
}

func (g *gatewayQueryBuilderNop) GuildUpdateChan(_ chan *GuildUpdate, _ ...chan *GuildUpdate) {
	return
}

func (g *gatewayQueryBuilderNop) InteractionCreate(_ func(Session, *InteractionCreate), _ ...func(Session, *InteractionCreate)) {
	return
}

func (g *gatewayQueryBuilderNop) InteractionCreateChan(_ chan *InteractionCreate, _ ...chan *InteractionCreate) {
	return
}

func (g *gatewayQueryBuilderNop) InviteCreate(_ func(Session, *InviteCreate), _ ...func(Session, *InviteCreate)) {
	return
}

func (g *gatewayQueryBuilderNop) InviteCreateChan(_ chan *InviteCreate, _ ...chan *InviteCreate) {
	return
}

func (g *gatewayQueryBuilderNop) InviteDelete(_ func(Session, *InviteDelete), _ ...func(Session, *InviteDelete)) {
	return
}

func (g *gatewayQueryBuilderNop) InviteDeleteChan(_ chan *InviteDelete, _ ...chan *InviteDelete) {
	return
}

func (g *gatewayQueryBuilderNop) MessageCreate(_ func(Session, *MessageCreate), _ ...func(Session, *MessageCreate)) {
	return
}

func (g *gatewayQueryBuilderNop) MessageCreateChan(_ chan *MessageCreate, _ ...chan *MessageCreate) {
	return
}

func (g *gatewayQueryBuilderNop) MessageDelete(_ func(Session, *MessageDelete), _ ...func(Session, *MessageDelete)) {
	return
}

func (g *gatewayQueryBuilderNop) MessageDeleteBulk(_ func(Session, *MessageDeleteBulk), _ ...func(Session, *MessageDeleteBulk)) {
	return
}

func (g *gatewayQueryBuilderNop) MessageDeleteBulkChan(_ chan *MessageDeleteBulk, _ ...chan *MessageDeleteBulk) {
	return
}

func (g *gatewayQueryBuilderNop) MessageDeleteChan(_ chan *MessageDelete, _ ...chan *MessageDelete) {
	return
}

func (g *gatewayQueryBuilderNop) MessageReactionAdd(_ func(Session, *MessageReactionAdd), _ ...func(Session, *MessageReactionAdd)) {
	return
}

func (g *gatewayQueryBuilderNop) MessageReactionAddChan(_ chan *MessageReactionAdd, _ ...chan *MessageReactionAdd) {
	return
}

func (g *gatewayQueryBuilderNop) MessageReactionRemove(_ func(Session, *MessageReactionRemove), _ ...func(Session, *MessageReactionRemove)) {
	return
}

func (g *gatewayQueryBuilderNop) MessageReactionRemoveAll(_ func(Session, *MessageReactionRemoveAll), _ ...func(Session, *MessageReactionRemoveAll)) {
	return
}

func (g *gatewayQueryBuilderNop) MessageReactionRemoveAllChan(_ chan *MessageReactionRemoveAll, _ ...chan *MessageReactionRemoveAll) {
	return
}

func (g *gatewayQueryBuilderNop) MessageReactionRemoveChan(_ chan *MessageReactionRemove, _ ...chan *MessageReactionRemove) {
	return
}

func (g *gatewayQueryBuilderNop) MessageReactionRemoveEmoji(_ func(Session, *MessageReactionRemoveEmoji), _ ...func(Session, *MessageReactionRemoveEmoji)) {
	return
}

func (g *gatewayQueryBuilderNop) MessageReactionRemoveEmojiChan(_ chan *MessageReactionRemoveEmoji, _ ...chan *MessageReactionRemoveEmoji) {
	return
}

func (g *gatewayQueryBuilderNop) MessageUpdate(_ func(Session, *MessageUpdate), _ ...func(Session, *MessageUpdate)) {
	return
}

func (g *gatewayQueryBuilderNop) MessageUpdateChan(_ chan *MessageUpdate, _ ...chan *MessageUpdate) {
	return
}

func (g *gatewayQueryBuilderNop) PresenceUpdate(_ func(Session, *PresenceUpdate), _ ...func(Session, *PresenceUpdate)) {
	return
}

func (g *gatewayQueryBuilderNop) PresenceUpdateChan(_ chan *PresenceUpdate, _ ...chan *PresenceUpdate) {
	return
}

func (g *gatewayQueryBuilderNop) Ready(_ func(Session, *Ready), _ ...func(Session, *Ready)) {
	return
}

func (g *gatewayQueryBuilderNop) ReadyChan(_ chan *Ready, _ ...chan *Ready) {
	return
}

func (g *gatewayQueryBuilderNop) Resumed(_ func(Session, *Resumed), _ ...func(Session, *Resumed)) {
	return
}

func (g *gatewayQueryBuilderNop) ResumedChan(_ chan *Resumed, _ ...chan *Resumed) {
	return
}

func (g *gatewayQueryBuilderNop) StayConnectedUntilInterrupted() error {
	return nil
}

func (g *gatewayQueryBuilderNop) ThreadCreate(_ func(Session, *ThreadCreate), _ ...func(Session, *ThreadCreate)) {
	return
}

func (g *gatewayQueryBuilderNop) ThreadCreateChan(_ chan *ThreadCreate, _ ...chan *ThreadCreate) {
	return
}

func (g *gatewayQueryBuilderNop) ThreadDelete(_ func(Session, *ThreadDelete), _ ...func(Session, *ThreadDelete)) {
	return
}

func (g *gatewayQueryBuilderNop) ThreadDeleteChan(_ chan *ThreadDelete, _ ...chan *ThreadDelete) {
	return
}

func (g *gatewayQueryBuilderNop) ThreadListSync(_ func(Session, *ThreadListSync), _ ...func(Session, *ThreadListSync)) {
	return
}

func (g *gatewayQueryBuilderNop) ThreadListSyncChan(_ chan *ThreadListSync, _ ...chan *ThreadListSync) {
	return
}

func (g *gatewayQueryBuilderNop) ThreadMembersUpdate(_ func(Session, *ThreadMembersUpdate), _ ...func(Session, *ThreadMembersUpdate)) {
	return
}

func (g *gatewayQueryBuilderNop) ThreadMembersUpdateChan(_ chan *ThreadMembersUpdate, _ ...chan *ThreadMembersUpdate) {
	return
}

func (g *gatewayQueryBuilderNop) ThreadMemberUpdate(_ func(Session, *ThreadMemberUpdate), _ ...func(Session, *ThreadMemberUpdate)) {
	return
}

func (g *gatewayQueryBuilderNop) ThreadMemberUpdateChan(_ chan *ThreadMemberUpdate, _ ...chan *ThreadMemberUpdate) {
	return
}

func (g *gatewayQueryBuilderNop) ThreadUpdate(_ func(Session, *ThreadUpdate), _ ...func(Session, *ThreadUpdate)) {
	return
}

func (g *gatewayQueryBuilderNop) ThreadUpdateChan(_ chan *ThreadUpdate, _ ...chan *ThreadUpdate) {
	return
}

func (g *gatewayQueryBuilderNop) TypingStart(_ func(Session, *TypingStart), _ ...func(Session, *TypingStart)) {
	return
}

func (g *gatewayQueryBuilderNop) TypingStartChan(_ chan *TypingStart, _ ...chan *TypingStart) {
	return
}

func (g *gatewayQueryBuilderNop) UserUpdate(_ func(Session, *UserUpdate), _ ...func(Session, *UserUpdate)) {
	return
}

func (g *gatewayQueryBuilderNop) UserUpdateChan(_ chan *UserUpdate, _ ...chan *UserUpdate) {
	return
}

func (g *gatewayQueryBuilderNop) VoiceServerUpdate(_ func(Session, *VoiceServerUpdate), _ ...func(Session, *VoiceServerUpdate)) {
	return
}

func (g *gatewayQueryBuilderNop) VoiceServerUpdateChan(_ chan *VoiceServerUpdate, _ ...chan *VoiceServerUpdate) {
	return
}

func (g *gatewayQueryBuilderNop) VoiceStateUpdate(_ func(Session, *VoiceStateUpdate), _ ...func(Session, *VoiceStateUpdate)) {
	return
}

func (g *gatewayQueryBuilderNop) VoiceStateUpdateChan(_ chan *VoiceStateUpdate, _ ...chan *VoiceStateUpdate) {
	return
}

func (g *gatewayQueryBuilderNop) WebhooksUpdate(_ func(Session, *WebhooksUpdate), _ ...func(Session, *WebhooksUpdate)) {
	return
}

func (g *gatewayQueryBuilderNop) WebhooksUpdateChan(_ chan *WebhooksUpdate, _ ...chan *WebhooksUpdate) {
	return
}

func (g *gatewayQueryBuilderNop) WithCtrl(_ HandlerCtrl) SocketHandlerRegistrator {
	return nil
}

func (g *gatewayQueryBuilderNop) WithMiddleware(_ func(interface{}) interface{}, _ ...func(interface{}) interface{}) SocketHandlerRegistrator {
	return nil
}

type guildEmojiQueryBuilderNop struct {
	Ctx       context.Context
	Flags     Flag
	ChannelID Snowflake
	GuildID   Snowflake
	UserID    Snowflake
}

var _ GuildEmojiQueryBuilder = &guildEmojiQueryBuilderNop{}

func (g guildEmojiQueryBuilderNop) WithContext(ctx context.Context) GuildEmojiQueryBuilder {
	g.Ctx = ctx
	return &g
}

func (g guildEmojiQueryBuilderNop) WithFlags(flags ...Flag) GuildEmojiQueryBuilder {
	g.Flags = mergeFlags(flags)
	return &g
}

func (g *guildEmojiQueryBuilderNop) Delete() error {
	return nil
}

func (g *guildEmojiQueryBuilderNop) Get() (*Emoji, error) {
	return nil, nil
}

func (g *guildEmojiQueryBuilderNop) Update(_ *UpdateEmoji) (*Emoji, error) {
	return nil, nil
}

func (g *guildEmojiQueryBuilderNop) UpdateBuilder() UpdateGuildEmojiBuilder {
	return nil
}

type guildMemberQueryBuilderNop struct {
	Ctx       context.Context
	Flags     Flag
	ChannelID Snowflake
	GuildID   Snowflake
	UserID    Snowflake
}

var _ GuildMemberQueryBuilder = &guildMemberQueryBuilderNop{}

func (g guildMemberQueryBuilderNop) WithContext(ctx context.Context) GuildMemberQueryBuilder {
	g.Ctx = ctx
	return &g
}

func (g guildMemberQueryBuilderNop) WithFlags(flags ...Flag) GuildMemberQueryBuilder {
	g.Flags = mergeFlags(flags)
	return &g
}

func (g *guildMemberQueryBuilderNop) AddRole(_ Snowflake) error {
	return nil
}

func (g *guildMemberQueryBuilderNop) Ban(_ *BanMember) error {
	return nil
}

func (g *guildMemberQueryBuilderNop) Get() (*Member, error) {
	return nil, nil
}

func (g *guildMemberQueryBuilderNop) GetPermissions() (PermissionBit, error) {
	return 0, nil
}

func (g *guildMemberQueryBuilderNop) Kick(_ string) error {
	return nil
}

func (g *guildMemberQueryBuilderNop) RemoveRole(_ Snowflake) error {
	return nil
}

func (g *guildMemberQueryBuilderNop) Update(_ *UpdateMember) (*Member, error) {
	return nil, nil
}

func (g *guildMemberQueryBuilderNop) UpdateBuilder() UpdateGuildMemberBuilder {
	return nil
}

type guildQueryBuilderNop struct {
	Ctx       context.Context
	Flags     Flag
	ChannelID Snowflake
	GuildID   Snowflake
	UserID    Snowflake
}

var _ GuildQueryBuilder = &guildQueryBuilderNop{}

func (g guildQueryBuilderNop) WithContext(ctx context.Context) GuildQueryBuilder {
	g.Ctx = ctx
	return &g
}

func (g guildQueryBuilderNop) WithFlags(flags ...Flag) GuildQueryBuilder {
	g.Flags = mergeFlags(flags)
	return &g
}

func (g *guildQueryBuilderNop) CreateChannel(_ string, _ *CreateGuildChannel) (*Channel, error) {
	return nil, nil
}

func (g *guildQueryBuilderNop) CreateEmoji(_ *CreateGuildEmoji) (*Emoji, error) {
	return nil, nil
}

func (g *guildQueryBuilderNop) CreateIntegration(_ *CreateGuildIntegration) error {
	return nil
}

func (g *guildQueryBuilderNop) CreateMember(_ Snowflake, _ string, _ *AddGuildMember) (*Member, error) {
	return nil, nil
}

func (g *guildQueryBuilderNop) CreateRole(_ *CreateGuildRole) (*Role, error) {
	return nil, nil
}

func (g *guildQueryBuilderNop) Delete() error {
	return nil
}

func (g *guildQueryBuilderNop) DeleteIntegration(_ Snowflake) error {
	return nil
}

func (g *guildQueryBuilderNop) DisconnectVoiceParticipant(_ Snowflake) error {
	return nil
}

func (g *guildQueryBuilderNop) Emoji(_ Snowflake) GuildEmojiQueryBuilder {
	return nil
}

func (g *guildQueryBuilderNop) EstimatePruneMembersCount(_ int) (int, error) {
	return 0, nil
}

func (g *guildQueryBuilderNop) Get() (*Guild, error) {
	return nil, nil
}

func (g *guildQueryBuilderNop) GetActiveThreads() (*ActiveGuildThreads, error) {
	return nil, nil
}

func (g *guildQueryBuilderNop) GetAuditLogs(_ *GetAuditLogs) (*AuditLog, error) {
	return nil, nil
}

func (g *guildQueryBuilderNop) GetBan(_ Snowflake) (*Ban, error) {
	return nil, nil
}

func (g *guildQueryBuilderNop) GetBans() ([]*Ban, error) {
	return nil, nil
}

func (g *guildQueryBuilderNop) GetChannels() ([]*Channel, error) {
	return nil, nil
}

func (g *guildQueryBuilderNop) GetEmbed() (*GuildWidget, error) {
	return nil, nil
}

func (g *guildQueryBuilderNop) GetEmojis() ([]*Emoji, error) {
	return nil, nil
}

func (g *guildQueryBuilderNop) GetIntegrations() ([]*Integration, error) {
	return nil, nil
}

func (g *guildQueryBuilderNop) GetInvites() ([]*Invite, error) {
	return nil, nil
}

func (g *guildQueryBuilderNop) GetMembers(_ *GetMembers) ([]*Member, error) {
	return nil, nil
}

func (g *guildQueryBuilderNop) GetPruneMembersCount(_ *GetPruneMembersCount) (int, error) {
	return 0, nil
}

func (g *guildQueryBuilderNop) GetRoles() ([]*Role, error) {
	return nil, nil
}

func (g *guildQueryBuilderNop) GetVanityURL() (*Invite, error) {
	return nil, nil
}

func (g *guildQueryBuilderNop) GetVoiceRegions() ([]*VoiceRegion, error) {
	return nil, nil
}

func (g *guildQueryBuilderNop) GetWebhooks() ([]*Webhook, error) {
	return nil, nil
}

func (g *guildQueryBuilderNop) GetWidget() (*GuildWidget, error) {
	return nil, nil
}

func (g *guildQueryBuilderNop) KickVoiceParticipant(_ Snowflake) error {
	return nil
}

func (g *guildQueryBuilderNop) Leave() error {
	return nil
}

func (g *guildQueryBuilderNop) Member(_ Snowflake) GuildMemberQueryBuilder {
	return nil
}

func (g *guildQueryBuilderNop) PruneMembers(_ *PruneMembers) (int, error) {
	return 0, nil
}

func (g *guildQueryBuilderNop) Role(_ Snowflake) GuildRoleQueryBuilder {
	return nil
}

func (g *guildQueryBuilderNop) SetCurrentUserNick(_ string) (string, error) {
	return "", nil
}

func (g *guildQueryBuilderNop) SyncIntegration(_ Snowflake) error {
	return nil
}

func (g *guildQueryBuilderNop) UnbanUser(_ Snowflake, _ string) error {
	return nil
}

func (g *guildQueryBuilderNop) Update(_ *UpdateGuild) (*Guild, error) {
	return nil, nil
}

func (g *guildQueryBuilderNop) UpdateBuilder() UpdateGuildBuilder {
	return nil
}

func (g *guildQueryBuilderNop) UpdateChannelPositions(_ []UpdateGuildChannelPositions) error {
	return nil
}

func (g *guildQueryBuilderNop) UpdateEmbedBuilder() UpdateGuildEmbedBuilder {
	return nil
}

func (g *guildQueryBuilderNop) UpdateIntegration(_ Snowflake, _ *UpdateGuildIntegration) error {
	return nil
}

func (g *guildQueryBuilderNop) UpdateRolePositions(_ []UpdateGuildRolePositions) ([]*Role, error) {
	return nil, nil
}

func (g *guildQueryBuilderNop) UpdateWidget(_ *UpdateGuildWidget) (*GuildWidget, error) {
	return nil, nil
}

func (g *guildQueryBuilderNop) VoiceChannel(_ Snowflake) VoiceChannelQueryBuilder {
	return nil
}

type guildRoleQueryBuilderNop struct {
	Ctx       context.Context
	Flags     Flag
	ChannelID Snowflake
	GuildID   Snowflake
	UserID    Snowflake
}

var _ GuildRoleQueryBuilder = &guildRoleQueryBuilderNop{}

func (g guildRoleQueryBuilderNop) WithContext(ctx context.Context) GuildRoleQueryBuilder {
	g.Ctx = ctx
	return &g
}

func (g guildRoleQueryBuilderNop) WithFlags(flags ...Flag) GuildRoleQueryBuilder {
	g.Flags = mergeFlags(flags)
	return &g
}

func (g *guildRoleQueryBuilderNop) Delete() error {
	return nil
}

func (g *guildRoleQueryBuilderNop) Update(_ *UpdateRole) (*Role, error) {
	return nil, nil
}

func (g *guildRoleQueryBuilderNop) UpdateBuilder() UpdateGuildRoleBuilder {
	return nil
}

type inviteQueryBuilderNop struct {
	Ctx       context.Context
	Flags     Flag
	ChannelID Snowflake
	GuildID   Snowflake
	UserID    Snowflake
}

var _ InviteQueryBuilder = &inviteQueryBuilderNop{}

func (i inviteQueryBuilderNop) WithContext(ctx context.Context) InviteQueryBuilder {
	i.Ctx = ctx
	return &i
}

func (i inviteQueryBuilderNop) WithFlags(flags ...Flag) InviteQueryBuilder {
	i.Flags = mergeFlags(flags)
	return &i
}

func (i *inviteQueryBuilderNop) Delete() (*Invite, error) {
	return nil, nil
}

func (i *inviteQueryBuilderNop) Get(_ bool) (*Invite, error) {
	return nil, nil
}

type messageQueryBuilderNop struct {
	Ctx       context.Context
	Flags     Flag
	ChannelID Snowflake
	GuildID   Snowflake
	UserID    Snowflake
}

var _ MessageQueryBuilder = &messageQueryBuilderNop{}

func (m messageQueryBuilderNop) WithContext(ctx context.Context) MessageQueryBuilder {
	m.Ctx = ctx
	return &m
}

func (m messageQueryBuilderNop) WithFlags(flags ...Flag) MessageQueryBuilder {
	m.Flags = mergeFlags(flags)
	return &m
}

func (m *messageQueryBuilderNop) CreateThread(_ *CreateThread) (*Channel, error) {
	return nil, nil
}

func (m *messageQueryBuilderNop) CrossPost() (*Message, error) {
	return nil, nil
}

func (m *messageQueryBuilderNop) Delete() error {
	return nil
}

func (m *messageQueryBuilderNop) DeleteAllReactions() error {
	return nil
}

func (m *messageQueryBuilderNop) Get() (*Message, error) {
	return nil, nil
}

func (m *messageQueryBuilderNop) Pin() error {
	return nil
}

func (m *messageQueryBuilderNop) Reaction(_ interface{}) ReactionQueryBuilder {
	return nil
}

func (m *messageQueryBuilderNop) SetContent(_ string) (*Message, error) {
	return nil, nil
}

func (m *messageQueryBuilderNop) SetEmbed(_ *Embed) (*Message, error) {
	return nil, nil
}

func (m *messageQueryBuilderNop) Unpin() error {
	return nil
}

func (m *messageQueryBuilderNop) Update(_ *UpdateMessage) (*Message, error) {
	return nil, nil
}

func (m *messageQueryBuilderNop) UpdateBuilder() UpdateMessageBuilder {
	return nil
}

type reactionQueryBuilderNop struct {
	Ctx       context.Context
	Flags     Flag
	ChannelID Snowflake
	GuildID   Snowflake
	UserID    Snowflake
}

var _ ReactionQueryBuilder = &reactionQueryBuilderNop{}

func (r reactionQueryBuilderNop) WithContext(ctx context.Context) ReactionQueryBuilder {
	r.Ctx = ctx
	return &r
}

func (r reactionQueryBuilderNop) WithFlags(flags ...Flag) ReactionQueryBuilder {
	r.Flags = mergeFlags(flags)
	return &r
}

func (r *reactionQueryBuilderNop) Create() error {
	return nil
}

func (r *reactionQueryBuilderNop) DeleteOwn() error {
	return nil
}

func (r *reactionQueryBuilderNop) DeleteUser(_ Snowflake) error {
	return nil
}

func (r *reactionQueryBuilderNop) Get(_ URLQueryStringer) ([]*User, error) {
	return nil, nil
}

type userQueryBuilderNop struct {
	Ctx       context.Context
	Flags     Flag
	ChannelID Snowflake
	GuildID   Snowflake
	UserID    Snowflake
}

var _ UserQueryBuilder = &userQueryBuilderNop{}

func (u userQueryBuilderNop) WithContext(ctx context.Context) UserQueryBuilder {
	u.Ctx = ctx
	return &u
}

func (u userQueryBuilderNop) WithFlags(flags ...Flag) UserQueryBuilder {
	u.Flags = mergeFlags(flags)
	return &u
}

func (u *userQueryBuilderNop) CreateDM() (*Channel, error) {
	return nil, nil
}

func (u *userQueryBuilderNop) Get() (*User, error) {
	return nil, nil
}

type voiceChannelQueryBuilderNop struct {
	Ctx       context.Context
	Flags     Flag
	ChannelID Snowflake
	GuildID   Snowflake
	UserID    Snowflake
}

var _ VoiceChannelQueryBuilder = &voiceChannelQueryBuilderNop{}

func (v voiceChannelQueryBuilderNop) WithContext(ctx context.Context) VoiceChannelQueryBuilder {
	v.Ctx = ctx
	return &v
}

func (v voiceChannelQueryBuilderNop) WithFlags(flags ...Flag) VoiceChannelQueryBuilder {
	v.Flags = mergeFlags(flags)
	return &v
}

func (v *voiceChannelQueryBuilderNop) Connect(_ bool, _ bool) (VoiceConnection, error) {
	return nil, nil
}

func (v *voiceChannelQueryBuilderNop) CreateInvite(_ *CreateInvite) (*Invite, error) {
	return nil, nil
}

func (v *voiceChannelQueryBuilderNop) Delete() (*Channel, error) {
	return nil, nil
}

func (v *voiceChannelQueryBuilderNop) DeletePermission(_ Snowflake) error {
	return nil
}

func (v *voiceChannelQueryBuilderNop) Get() (*Channel, error) {
	return nil, nil
}

func (v *voiceChannelQueryBuilderNop) GetInvites() ([]*Invite, error) {
	return nil, nil
}

func (v *voiceChannelQueryBuilderNop) JoinManual(_ bool, _ bool) (*VoiceStateUpdate, *VoiceServerUpdate, error) {
	return nil, nil, nil
}

func (v *voiceChannelQueryBuilderNop) Update(_ *UpdateChannel) (*Channel, error) {
	return nil, nil
}

func (v *voiceChannelQueryBuilderNop) UpdateBuilder() UpdateChannelBuilder {
	return nil
}

func (v *voiceChannelQueryBuilderNop) UpdatePermissions(_ Snowflake, _ *UpdateChannelPermissions) error {
	return nil
}

type webhookQueryBuilderNop struct {
	Ctx       context.Context
	Flags     Flag
	ChannelID Snowflake
	GuildID   Snowflake
	UserID    Snowflake
}

var _ WebhookQueryBuilder = &webhookQueryBuilderNop{}

func (w webhookQueryBuilderNop) WithContext(ctx context.Context) WebhookQueryBuilder {
	w.Ctx = ctx
	return &w
}

func (w webhookQueryBuilderNop) WithFlags(flags ...Flag) WebhookQueryBuilder {
	w.Flags = mergeFlags(flags)
	return &w
}

func (w *webhookQueryBuilderNop) Delete() error {
	return nil
}

func (w *webhookQueryBuilderNop) Get() (*Webhook, error) {
	return nil, nil
}

func (w *webhookQueryBuilderNop) Update(_ *UpdateWebhook) (*Webhook, error) {
	return nil, nil
}

func (w *webhookQueryBuilderNop) UpdateBuilder() UpdateWebhookBuilder {
	return nil
}

func (w *webhookQueryBuilderNop) WithToken(_ string) WebhookWithTokenQueryBuilder {
	return nil
}

type webhookWithTokenQueryBuilderNop struct {
	Ctx       context.Context
	Flags     Flag
	ChannelID Snowflake
	GuildID   Snowflake
	UserID    Snowflake
}

var _ WebhookWithTokenQueryBuilder = &webhookWithTokenQueryBuilderNop{}

func (w webhookWithTokenQueryBuilderNop) WithContext(ctx context.Context) WebhookWithTokenQueryBuilder {
	w.Ctx = ctx
	return &w
}

func (w webhookWithTokenQueryBuilderNop) WithFlags(flags ...Flag) WebhookWithTokenQueryBuilder {
	w.Flags = mergeFlags(flags)
	return &w
}

func (w *webhookWithTokenQueryBuilderNop) Delete() error {
	return nil
}

func (w *webhookWithTokenQueryBuilderNop) DeleteMessage(_ Snowflake, _ *Snowflake) error {
	return nil
}

func (w *webhookWithTokenQueryBuilderNop) EditMessage(_ *ExecuteWebhook, _ Snowflake, _ *Snowflake) (*Message, error) {
	return nil, nil
}

func (w *webhookWithTokenQueryBuilderNop) Execute(_ *ExecuteWebhook, _ *bool, _ *Snowflake, _ string) (*Message, error) {
	return nil, nil
}

func (w *webhookWithTokenQueryBuilderNop) ExecuteGitHubWebhook(_ *ExecuteWebhook, _ *bool, _ *Snowflake) (*Message, error) {
	return nil, nil
}

func (w *webhookWithTokenQueryBuilderNop) ExecuteSlackWebhook(_ *ExecuteWebhook, _ *bool, _ *Snowflake) (*Message, error) {
	return nil, nil
}

func (w *webhookWithTokenQueryBuilderNop) Get() (*Webhook, error) {
	return nil, nil
}

func (w *webhookWithTokenQueryBuilderNop) GetMessage(_ Snowflake, _ *Snowflake) (*Message, error) {
	return nil, nil
}

func (w *webhookWithTokenQueryBuilderNop) Update(_ *UpdateWebhook) (*Webhook, error) {
	return nil, nil
}

func (w *webhookWithTokenQueryBuilderNop) UpdateBuilder() UpdateWebhookBuilder {
	return nil
}
