package disgord

import (
	"fmt"
	"sort"
	"strings"
)

//////////////////////////////////////////////////////
//
// demultiplexer
//
//////////////////////////////////////////////////////

func Sort(v interface{}, fs ...Flag) {
	if v == nil {
		return
	}

	flags := mergeFlags(fs)
	if (flags & SortByID) > 0 {
		sortByID(v, flags)
	} else if (flags & SortByGuildID) > 0 {
		sortByGuildID(v, flags)
	} else if (flags & SortByChannelID) > 0 {
		sortByChannelID(v, flags)
	} else if (flags & SortByName) > 0 {
		sortByName(v, flags)
	} else if (flags & SortByHoist) > 0 {
		sortByHoist(v, flags)
	} else if list, ok := v.(sort.Interface); ok {
		if (flags & OrderDescending) > 0 {
			sort.Sort(sort.Reverse(list))
		} else {
			sort.Sort(list)
		}
	} else if list, ok := v.([]*Role); ok {
		if (flags & OrderDescending) > 0 {
			sort.Sort(sort.Reverse(roles(list)))
		} else {
			sort.Sort(roles(list))
		}
	} else if list, ok := v.(*[]*Role); ok {
		if (flags & OrderDescending) > 0 {
			sort.Sort(sort.Reverse(roles(*list)))
		} else {
			sort.Sort(roles(*list))
		}
	} else {
		panic("type is missing sort.Interface implementation")
	}
}

func derefSliceP(v interface{}) (s interface{}) {
	switch t := v.(type) {
	case *[]*Application:
		s = *t
	case *[]*ApplicationCommand:
		s = *t
	case *[]*ApplicationCommandDataOption:
		s = *t
	case *[]*ApplicationCommandOption:
		s = *t
	case *[]*ApplicationCommandOptionChoice:
		s = *t
	case *[]*ApplicationCommandPermissions:
		s = *t
	case *[]*CreateApplicationCommand:
		s = *t
	case *[]*GuildApplicationCommandPermissions:
		s = *t
	case *[]*UpdateApplicationCommand:
		s = *t
	case *[]*AuditLog:
		s = *t
	case *[]*AuditLogChanges:
		s = *t
	case *[]*AuditLogEntry:
		s = *t
	case *[]*AuditLogOption:
		s = *t
	case *[]*BasicCache:
		s = *t
	case *[]*AllowedMentions:
		s = *t
	case *[]*ArchivedThreads:
		s = *t
	case *[]*Attachment:
		s = *t
	case *[]*Channel:
		s = *t
	case *[]*CreateInvite:
		s = *t
	case *[]*CreateMessage:
		s = *t
	case *[]*CreateMessageFile:
		s = *t
	case *[]*CreateThreadWithoutMessage:
		s = *t
	case *[]*CreateWebhook:
		s = *t
	case *[]*DeleteMessages:
		s = *t
	case *[]*GetArchivedThreads:
		s = *t
	case *[]*GetMessages:
		s = *t
	case *[]*GroupDMParticipant:
		s = *t
	case *[]*PartialChannel:
		s = *t
	case *[]*PermissionOverwrite:
		s = *t
	case *[]*UpdateChannel:
		s = *t
	case *[]*UpdateChannelPermissions:
		s = *t
	case *[]*Client:
		s = *t
	case *[]*Config:
		s = *t
	case *[]*InteractionApplicationCommandCallbackData:
		s = *t
	case *[]*InteractionResponse:
		s = *t
	case *[]*ErrorEmptyValue:
		s = *t
	case *[]*ErrorMissingSnowflake:
		s = *t
	case *[]*Embed:
		s = *t
	case *[]*EmbedAuthor:
		s = *t
	case *[]*EmbedField:
		s = *t
	case *[]*EmbedFooter:
		s = *t
	case *[]*EmbedImage:
		s = *t
	case *[]*EmbedProvider:
		s = *t
	case *[]*EmbedThumbnail:
		s = *t
	case *[]*EmbedVideo:
		s = *t
	case *[]*Emoji:
		s = *t
	case *[]*UpdateEmoji:
		s = *t
	case *[]*ChannelCreate:
		s = *t
	case *[]*ChannelDelete:
		s = *t
	case *[]*ChannelPinsUpdate:
		s = *t
	case *[]*ChannelUpdate:
		s = *t
	case *[]*GuildBanAdd:
		s = *t
	case *[]*GuildBanRemove:
		s = *t
	case *[]*GuildCreate:
		s = *t
	case *[]*GuildDelete:
		s = *t
	case *[]*GuildEmojisUpdate:
		s = *t
	case *[]*GuildIntegrationsUpdate:
		s = *t
	case *[]*GuildMemberAdd:
		s = *t
	case *[]*GuildMemberRemove:
		s = *t
	case *[]*GuildMemberUpdate:
		s = *t
	case *[]*GuildMembersChunk:
		s = *t
	case *[]*GuildRoleCreate:
		s = *t
	case *[]*GuildRoleDelete:
		s = *t
	case *[]*GuildRoleUpdate:
		s = *t
	case *[]*GuildScheduledEventCreate:
		s = *t
	case *[]*GuildScheduledEventDelete:
		s = *t
	case *[]*GuildScheduledEventUpdate:
		s = *t
	case *[]*GuildScheduledEventUserAdd:
		s = *t
	case *[]*GuildScheduledEventUserRemove:
		s = *t
	case *[]*GuildStickersUpdate:
		s = *t
	case *[]*GuildUpdate:
		s = *t
	case *[]*InteractionCreate:
		s = *t
	case *[]*InviteCreate:
		s = *t
	case *[]*InviteDelete:
		s = *t
	case *[]*MessageCreate:
		s = *t
	case *[]*MessageDelete:
		s = *t
	case *[]*MessageDeleteBulk:
		s = *t
	case *[]*MessageReactionAdd:
		s = *t
	case *[]*MessageReactionRemove:
		s = *t
	case *[]*MessageReactionRemoveAll:
		s = *t
	case *[]*MessageReactionRemoveEmoji:
		s = *t
	case *[]*MessageUpdate:
		s = *t
	case *[]*PresenceUpdate:
		s = *t
	case *[]*Ready:
		s = *t
	case *[]*Resumed:
		s = *t
	case *[]*ThreadCreate:
		s = *t
	case *[]*ThreadDelete:
		s = *t
	case *[]*ThreadListSync:
		s = *t
	case *[]*ThreadMemberUpdate:
		s = *t
	case *[]*ThreadMembersUpdate:
		s = *t
	case *[]*ThreadUpdate:
		s = *t
	case *[]*TypingStart:
		s = *t
	case *[]*UserUpdate:
		s = *t
	case *[]*VoiceServerUpdate:
		s = *t
	case *[]*VoiceStateUpdate:
		s = *t
	case *[]*WebhooksUpdate:
		s = *t
	case *[]*ActiveGuildThreads:
		s = *t
	case *[]*AddGuildMember:
		s = *t
	case *[]*Ban:
		s = *t
	case *[]*BanMember:
		s = *t
	case *[]*CreateGuild:
		s = *t
	case *[]*CreateGuildChannel:
		s = *t
	case *[]*CreateGuildEmoji:
		s = *t
	case *[]*CreateGuildIntegration:
		s = *t
	case *[]*CreateGuildRole:
		s = *t
	case *[]*GetAuditLogs:
		s = *t
	case *[]*GetMembers:
		s = *t
	case *[]*GetPruneMembersCount:
		s = *t
	case *[]*Guild:
		s = *t
	case *[]*GuildUnavailable:
		s = *t
	case *[]*GuildWidget:
		s = *t
	case *[]*Integration:
		s = *t
	case *[]*IntegrationAccount:
		s = *t
	case *[]*Member:
		s = *t
	case *[]*PartialBan:
		s = *t
	case *[]*PruneMembers:
		s = *t
	case *[]*UpdateGuild:
		s = *t
	case *[]*UpdateGuildChannelPositions:
		s = *t
	case *[]*UpdateGuildIntegration:
		s = *t
	case *[]*UpdateGuildRolePositions:
		s = *t
	case *[]*UpdateGuildWidget:
		s = *t
	case *[]*GuildScheduledEvent:
		s = *t
	case *[]*ApplicationCommandInteractionData:
		s = *t
	case *[]*ApplicationCommandInteractionDataResolved:
		s = *t
	case *[]*CreateInteractionResponse:
		s = *t
	case *[]*CreateInteractionResponseData:
		s = *t
	case *[]*MessageInteraction:
		s = *t
	case *[]*Invite:
		s = *t
	case *[]*InviteMetadata:
		s = *t
	case *[]*UpdateMember:
		s = *t
	case *[]*CreateThread:
		s = *t
	case *[]*MentionChannel:
		s = *t
	case *[]*Message:
		s = *t
	case *[]*MessageActivity:
		s = *t
	case *[]*MessageApplication:
		s = *t
	case *[]*MessageComponent:
		s = *t
	case *[]*MessageReference:
		s = *t
	case *[]*MessageSticker:
		s = *t
	case *[]*SelectMenuOption:
		s = *t
	case *[]*StickerItem:
		s = *t
	case *[]*UpdateMessage:
		s = *t
	case *[]*GetReactionURL:
		s = *t
	case *[]*Reaction:
		s = *t
	case *[]*Ctrl:
		s = *t
	case *[]*RESTBuilder:
		s = *t
	case *[]*Role:
		s = *t
	case *[]*UpdateRole:
		s = *t
	case *[]*ErrorUnsupportedType:
		s = *t
	case *[]*Time:
		s = *t
	case *[]*ThreadMember:
		s = *t
	case *[]*ThreadMetadata:
		s = *t
	case *[]*Activity:
		s = *t
	case *[]*ActivityAssets:
		s = *t
	case *[]*ActivityEmoji:
		s = *t
	case *[]*ActivityParty:
		s = *t
	case *[]*ActivitySecrets:
		s = *t
	case *[]*ActivityTimestamp:
		s = *t
	case *[]*ClientStatus:
		s = *t
	case *[]*CreateGroupDM:
		s = *t
	case *[]*GetCurrentUserGuilds:
		s = *t
	case *[]*UpdateUser:
		s = *t
	case *[]*User:
		s = *t
	case *[]*UserConnection:
		s = *t
	case *[]*UserPresence:
		s = *t
	case *[]*VoiceRegion:
		s = *t
	case *[]*VoiceState:
		s = *t
	case *[]*ExecuteWebhook:
		s = *t
	case *[]*UpdateWebhook:
		s = *t
	case *[]*Webhook:
		s = *t
	default:
		s = t
	}

	return s
}
func sortByID(v interface{}, flags Flag) {
	var descending bool
	if (flags & OrderDescending) > 0 {
		descending = true
	}

	v = derefSliceP(v)

	var less func(i, j int) bool
	switch s := v.(type) {
	case []*Application:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*ApplicationCommand:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*ApplicationCommandPermissions:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*GuildApplicationCommandPermissions:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*AuditLogEntry:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*AuditLogOption:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*Attachment:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*Channel:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*PartialChannel:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*PermissionOverwrite:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*Emoji:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*InteractionCreate:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*ThreadMembersUpdate:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*CreateGuildIntegration:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*Guild:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*GuildUnavailable:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*Integration:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*IntegrationAccount:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*UpdateGuildChannelPositions:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*UpdateGuildRolePositions:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*GuildScheduledEvent:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*ApplicationCommandInteractionData:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*MessageInteraction:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*MentionChannel:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*Message:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*MessageApplication:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*MessageSticker:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*StickerItem:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*Role:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*ActivityEmoji:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*ActivityParty:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*User:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*UserConnection:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*VoiceRegion:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	case []*Webhook:
		if descending {
			less = func(i, j int) bool { return s[i].ID > s[j].ID }
		} else {
			less = func(i, j int) bool { return s[i].ID < s[j].ID }
		}
	default:
		panic(fmt.Sprintf("type %+v does not support sorting", s))
	}

	sort.Slice(v, less)
}
func sortByGuildID(v interface{}, flags Flag) {
	var descending bool
	if (flags & OrderDescending) > 0 {
		descending = true
	}

	v = derefSliceP(v)

	var less func(i, j int) bool
	switch s := v.(type) {
	case []*ApplicationCommand:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*GuildApplicationCommandPermissions:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*Channel:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*ChannelPinsUpdate:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*GuildBanAdd:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*GuildBanRemove:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*GuildEmojisUpdate:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*GuildIntegrationsUpdate:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*GuildMemberRemove:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*GuildMembersChunk:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*GuildRoleCreate:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*GuildRoleDelete:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*GuildRoleUpdate:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*GuildScheduledEventUserAdd:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*GuildScheduledEventUserRemove:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*GuildStickersUpdate:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*InteractionCreate:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*InviteCreate:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*InviteDelete:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*MessageDelete:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*MessageReactionRemoveEmoji:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*PresenceUpdate:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*ThreadListSync:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*ThreadMembersUpdate:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*TypingStart:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*VoiceServerUpdate:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*WebhooksUpdate:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*Member:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*GuildScheduledEvent:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*MentionChannel:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*Message:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*MessageReference:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*UserPresence:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*VoiceState:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	case []*Webhook:
		if descending {
			less = func(i, j int) bool { return s[i].GuildID > s[j].GuildID }
		} else {
			less = func(i, j int) bool { return s[i].GuildID < s[j].GuildID }
		}
	default:
		panic(fmt.Sprintf("type %+v does not support sorting", s))
	}

	sort.Slice(v, less)
}
func sortByChannelID(v interface{}, flags Flag) {
	var descending bool
	if (flags & OrderDescending) > 0 {
		descending = true
	}

	v = derefSliceP(v)

	var less func(i, j int) bool
	switch s := v.(type) {
	case []*AuditLogOption:
		if descending {
			less = func(i, j int) bool { return s[i].ChannelID > s[j].ChannelID }
		} else {
			less = func(i, j int) bool { return s[i].ChannelID < s[j].ChannelID }
		}
	case []*ChannelPinsUpdate:
		if descending {
			less = func(i, j int) bool { return s[i].ChannelID > s[j].ChannelID }
		} else {
			less = func(i, j int) bool { return s[i].ChannelID < s[j].ChannelID }
		}
	case []*InteractionCreate:
		if descending {
			less = func(i, j int) bool { return s[i].ChannelID > s[j].ChannelID }
		} else {
			less = func(i, j int) bool { return s[i].ChannelID < s[j].ChannelID }
		}
	case []*InviteCreate:
		if descending {
			less = func(i, j int) bool { return s[i].ChannelID > s[j].ChannelID }
		} else {
			less = func(i, j int) bool { return s[i].ChannelID < s[j].ChannelID }
		}
	case []*InviteDelete:
		if descending {
			less = func(i, j int) bool { return s[i].ChannelID > s[j].ChannelID }
		} else {
			less = func(i, j int) bool { return s[i].ChannelID < s[j].ChannelID }
		}
	case []*MessageDelete:
		if descending {
			less = func(i, j int) bool { return s[i].ChannelID > s[j].ChannelID }
		} else {
			less = func(i, j int) bool { return s[i].ChannelID < s[j].ChannelID }
		}
	case []*MessageDeleteBulk:
		if descending {
			less = func(i, j int) bool { return s[i].ChannelID > s[j].ChannelID }
		} else {
			less = func(i, j int) bool { return s[i].ChannelID < s[j].ChannelID }
		}
	case []*MessageReactionAdd:
		if descending {
			less = func(i, j int) bool { return s[i].ChannelID > s[j].ChannelID }
		} else {
			less = func(i, j int) bool { return s[i].ChannelID < s[j].ChannelID }
		}
	case []*MessageReactionRemove:
		if descending {
			less = func(i, j int) bool { return s[i].ChannelID > s[j].ChannelID }
		} else {
			less = func(i, j int) bool { return s[i].ChannelID < s[j].ChannelID }
		}
	case []*MessageReactionRemoveAll:
		if descending {
			less = func(i, j int) bool { return s[i].ChannelID > s[j].ChannelID }
		} else {
			less = func(i, j int) bool { return s[i].ChannelID < s[j].ChannelID }
		}
	case []*MessageReactionRemoveEmoji:
		if descending {
			less = func(i, j int) bool { return s[i].ChannelID > s[j].ChannelID }
		} else {
			less = func(i, j int) bool { return s[i].ChannelID < s[j].ChannelID }
		}
	case []*TypingStart:
		if descending {
			less = func(i, j int) bool { return s[i].ChannelID > s[j].ChannelID }
		} else {
			less = func(i, j int) bool { return s[i].ChannelID < s[j].ChannelID }
		}
	case []*WebhooksUpdate:
		if descending {
			less = func(i, j int) bool { return s[i].ChannelID > s[j].ChannelID }
		} else {
			less = func(i, j int) bool { return s[i].ChannelID < s[j].ChannelID }
		}
	case []*GuildWidget:
		if descending {
			less = func(i, j int) bool { return s[i].ChannelID > s[j].ChannelID }
		} else {
			less = func(i, j int) bool { return s[i].ChannelID < s[j].ChannelID }
		}
	case []*GuildScheduledEvent:
		if descending {
			less = func(i, j int) bool { return s[i].ChannelID > s[j].ChannelID }
		} else {
			less = func(i, j int) bool { return s[i].ChannelID < s[j].ChannelID }
		}
	case []*Message:
		if descending {
			less = func(i, j int) bool { return s[i].ChannelID > s[j].ChannelID }
		} else {
			less = func(i, j int) bool { return s[i].ChannelID < s[j].ChannelID }
		}
	case []*MessageReference:
		if descending {
			less = func(i, j int) bool { return s[i].ChannelID > s[j].ChannelID }
		} else {
			less = func(i, j int) bool { return s[i].ChannelID < s[j].ChannelID }
		}
	case []*VoiceState:
		if descending {
			less = func(i, j int) bool { return s[i].ChannelID > s[j].ChannelID }
		} else {
			less = func(i, j int) bool { return s[i].ChannelID < s[j].ChannelID }
		}
	case []*Webhook:
		if descending {
			less = func(i, j int) bool { return s[i].ChannelID > s[j].ChannelID }
		} else {
			less = func(i, j int) bool { return s[i].ChannelID < s[j].ChannelID }
		}
	default:
		panic(fmt.Sprintf("type %+v does not support sorting", s))
	}

	sort.Slice(v, less)
}
func sortByName(v interface{}, flags Flag) {
	var descending bool
	if (flags & OrderDescending) > 0 {
		descending = true
	}

	v = derefSliceP(v)

	var less func(i, j int) bool
	switch s := v.(type) {
	case []*Application:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*ApplicationCommand:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*ApplicationCommandDataOption:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*ApplicationCommandOption:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*ApplicationCommandOptionChoice:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*CreateApplicationCommand:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*Channel:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*CreateThreadWithoutMessage:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*CreateWebhook:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*PartialChannel:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*EmbedAuthor:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*EmbedField:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*EmbedProvider:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*Emoji:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*CreateGuild:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*CreateGuildChannel:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*CreateGuildEmoji:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*CreateGuildRole:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*Guild:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*Integration:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*IntegrationAccount:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*GuildScheduledEvent:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*ApplicationCommandInteractionData:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*MessageInteraction:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*CreateThread:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*MentionChannel:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*MessageApplication:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*MessageSticker:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*StickerItem:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*Role:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*Activity:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*ActivityEmoji:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*UserConnection:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*VoiceRegion:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	case []*Webhook:
		if descending {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) > strings.ToLower(s[j].Name) }
		} else {
			less = func(i, j int) bool { return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name) }
		}
	default:
		panic(fmt.Sprintf("type %+v does not support sorting", s))
	}

	sort.Slice(v, less)
}
func sortByHoist(v interface{}, flags Flag) {
	var descending bool
	if (flags & OrderDescending) > 0 {
		descending = true
	}

	v = derefSliceP(v)

	var less func(i, j int) bool
	switch s := v.(type) {
	case []*CreateGuildRole:
		if descending {
			less = func(i, j int) bool { return s[i].Hoist && !s[j].Hoist }
		} else {
			less = func(i, j int) bool { return !s[i].Hoist && s[j].Hoist }
		}
	case []*Role:
		if descending {
			less = func(i, j int) bool { return s[i].Hoist && !s[j].Hoist }
		} else {
			less = func(i, j int) bool { return !s[i].Hoist && s[j].Hoist }
		}
	default:
		panic(fmt.Sprintf("type %+v does not support sorting", s))
	}

	sort.Slice(v, less)
}
