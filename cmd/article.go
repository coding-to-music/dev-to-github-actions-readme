package cmd

import (
	"context"
	"log/slog"
	"os"

	"github.com/coding-to-music/dev-to-github-actions-readme/handler/collector"
	"github.com/coding-to-music/dev-to-github-actions-readme/impl/article_service/forem"
	foremsrv "github.com/coding-to-music/dev-to-github-actions-readme/pkg/forem"
	"github.com/spf13/cobra"
)

func UpdateArticles(use string) *cobra.Command {
	var templateFilePath string
	var outputFilePath string
	var username string
	var limit int

	command := &cobra.Command{
		Use: use,
		Run: func(cmd *cobra.Command, args []string) {
			devToService := forem.NewService(foremsrv.NewService(foremsrv.DevToEndpoint))
			handler := collector.NewHandler(devToService)
			err := handler.Collect(context.Background(), username, limit, templateFilePath, outputFilePath)
			if err != nil {
				slog.Error(err.Error())
				os.Exit(1)
			}
			slog.Info("Updated!")
		},
	}

	command.Flags().StringVarP(&templateFilePath, "template-file", "f", "", "Readme template file path")
	command.Flags().StringVarP(&outputFilePath, "out-file", "o", "", "Output file path")
	command.Flags().StringVarP(&username, "username", "u", "", "Username")
	command.Flags().IntVar(&limit, "limit", 5, "Limit number of articles")
	err := command.MarkFlagRequired("username")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	err = command.MarkFlagRequired("template-file")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	err = command.MarkFlagRequired("out-file")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	return command
}
