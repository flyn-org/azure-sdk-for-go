//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azopenai_test

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestChatCompletions_extensions_bringYourOwnData(t *testing.T) {
	client := newAzureOpenAIClientForTest(t, azureOpenAI)

	resp, err := client.GetChatCompletions(context.Background(), azopenai.ChatCompletionsOptions{
		Messages: []azopenai.ChatMessage{
			{Content: to.Ptr("What does PR complete mean?"), Role: to.Ptr(azopenai.ChatRoleUser)},
		},
		MaxTokens: to.Ptr[int32](512),
		AzureExtensionsOptions: &azopenai.AzureChatExtensionOptions{
			Extensions: []azopenai.AzureChatExtensionConfiguration{
				{
					Type:       to.Ptr(azopenai.AzureChatExtensionTypeAzureCognitiveSearch),
					Parameters: azureOpenAI.Cognitive,
				},
			},
		},
		Deployment: "gpt-4",
	}, nil)
	require.NoError(t, err)

	// when you BYOD you get some extra content showing you metadata/info from the external
	// data source.
	msgContext := resp.Choices[0].Message.Context
	require.NotEmpty(t, msgContext.Messages[0].Content)
	require.Equal(t, azopenai.ChatRoleTool, *msgContext.Messages[0].Role)

	require.NotEmpty(t, *resp.Choices[0].Message.Content)
	require.Equal(t, azopenai.CompletionsFinishReasonStop, *resp.Choices[0].FinishReason)
}

func TestChatExtensionsStreaming_extensions_bringYourOwnData(t *testing.T) {
	client := newAzureOpenAIClientForTest(t, azureOpenAI)

	streamResp, err := client.GetChatCompletionsStream(context.Background(), azopenai.ChatCompletionsOptions{
		Messages: []azopenai.ChatMessage{
			{Content: to.Ptr("What does PR complete mean?"), Role: to.Ptr(azopenai.ChatRoleUser)},
		},
		MaxTokens: to.Ptr[int32](512),
		AzureExtensionsOptions: &azopenai.AzureChatExtensionOptions{
			Extensions: []azopenai.AzureChatExtensionConfiguration{
				{
					Type:       to.Ptr(azopenai.AzureChatExtensionTypeAzureCognitiveSearch),
					Parameters: azureOpenAI.Cognitive,
				},
			},
		},
		Deployment: "gpt-4",
	}, nil)

	require.NoError(t, err)
	defer streamResp.ChatCompletionsStream.Close()

	text := ""

	first := false

	for {
		event, err := streamResp.ChatCompletionsStream.Read()

		if errors.Is(err, io.EOF) {
			break
		}

		require.NoError(t, err)

		if first {
			// when you BYOD you get some extra content showing you metadata/info from the external
			// data source.
			first = false
			msgContext := event.Choices[0].Message.Context
			require.NotEmpty(t, msgContext.Messages[0].Content)
			require.Equal(t, azopenai.ChatRoleTool, *msgContext.Messages[0].Role)
		}

		for _, choice := range event.Choices {
			if choice.Delta != nil && choice.Delta.Content != nil {
				text += *choice.Delta.Content
			}
		}
	}

	require.NotEmpty(t, text)
}
