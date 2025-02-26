package basiccommands

import (
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/internal/database"
	"github.com/itsjustabavaar/oh-my-gosh/internal/models"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"sort"
	"strings"
	"time"
)

type HistoryCommand struct {
	Output string
}

type CommandSummary struct {
	Command    string
	Count      int
	LatestUsed time.Time
}

func (h *HistoryCommand) Handler() (string, *int, error) {
	commandsSummary := make([]CommandSummary, 0)
	var err error

	if vars.CurrentUser.Username == "" {
		commandsSummary, err = GetAnonymousCommandsSummary()
		if err != nil {
			return "", nil, err
		}
	} else {
		commandsSummary, err = GetUserCommandsSummary()
		if err != nil {
			return "", nil, err
		}
	}
	result := FormatCommandSummaryTable(commandsSummary)
	return result, nil, nil
}

func AddUserHistory(user *models.User, command string) error {
	history := models.CommandHistory{
		UserID:    user.ID,
		Command:   command,
		Timestamp: time.Now(),
	}
	return database.GetDB().Create(&history).Error
}

func AddAnonymousHistory(command string) {
	history := vars.History{
		Command:   command,
		Timestamp: time.Now(),
	}
	vars.AnonymousHistory = append(vars.AnonymousHistory, history)
}

func StoreCommandHistory(input string) error {
	if vars.CurrentUser.Username == "" {
		AddAnonymousHistory(input)
	} else {
		err := AddUserHistory(vars.CurrentUser, input)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetUserCommandsSummary() ([]CommandSummary, error) {
	var results []struct {
		Command    string
		Count      int
		LatestUsed time.Time
	}

	err := database.GetDB().Model(&models.CommandHistory{}).
		Select("command, COUNT(*) as count, MAX(timestamp) as latest_used").
		Where("user_id = ?", vars.CurrentUser.ID).
		Group("command").
		Order("count DESC, latest_used DESC").
		Find(&results).Error

	if err != nil {
		return nil, err
	}

	summaries := make([]CommandSummary, len(results))
	for i, result := range results {
		summaries[i] = CommandSummary{
			Command:    result.Command,
			Count:      result.Count,
			LatestUsed: result.LatestUsed,
		}
	}

	return summaries, nil
}

func GetAnonymousCommandsSummary() ([]CommandSummary, error) {
	summaryMap := make(map[string]*CommandSummary)

	for _, entry := range vars.AnonymousHistory {
		if entry.Command == "" {
			continue
		}

		summary, exists := summaryMap[entry.Command]
		if !exists {
			summaryMap[entry.Command] = &CommandSummary{
				Command:    entry.Command,
				Count:      1,
				LatestUsed: entry.Timestamp,
			}
		} else {
			summary.Count++
			if entry.Timestamp.After(summary.LatestUsed) {
				summary.LatestUsed = entry.Timestamp
			}
		}
	}

	summaries := make([]CommandSummary, 0, len(summaryMap))
	for _, summary := range summaryMap {
		summaries = append(summaries, *summary)
	}

	sort.Slice(summaries, func(i, j int) bool {
		if summaries[i].Count != summaries[j].Count {
			return summaries[i].Count > summaries[j].Count
		}
		return summaries[i].LatestUsed.After(summaries[j].LatestUsed)
	})

	return summaries, nil
}

func FormatCommandSummaryTable(summaries []CommandSummary) string {
	if len(summaries) == 0 {
		return "No commands found"
	}

	commandMaxLength := 0
	for _, summary := range summaries {
		if len(summary.Command) > commandMaxLength {
			commandMaxLength = len(summary.Command)
		}
	}

	var historyLines []string

	for _, summary := range summaries {
		var builder strings.Builder
		padding := commandMaxLength - len(summary.Command)
		builder.WriteString("| ")
		builder.WriteString(summary.Command)
		builder.WriteString(strings.Repeat(" ", padding))
		builder.WriteString(" | ")
		builder.WriteString(fmt.Sprintf("%d", summary.Count))
		builder.WriteString(" | ")
		historyLines = append(historyLines, builder.String())
	}
	return strings.Join(historyLines, "\n")
}
