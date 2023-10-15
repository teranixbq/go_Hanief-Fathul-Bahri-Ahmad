package usecase

import (
    "context"
    "fmt"
    "github.com/sashabaranov/go-openai"
    "openai/dto"
)

type LaptopRecomendedUsecase interface {
    LaptopRecomended(request dto.RequestRecomended, key string) (string, error)
}

type Usecase struct{}

func NewUseCase() LaptopRecomendedUsecase {
    return &Usecase{}
}

func (uc *Usecase) getCompletionMessages(ctx context.Context, client *openai.Client, messages []openai.ChatCompletionMessage, model string) (openai.ChatCompletionResponse, error) {
    if model == "" {
        model = openai.GPT3Dot5Turbo
    }

    resp, err := client.CreateChatCompletion(
        ctx,
        openai.ChatCompletionRequest{
            Model:    model,
            Messages: messages,
        },
    )
    return resp, err
}

func (us *Usecase) LaptopRecomended(request dto.RequestRecomended, key string) (string, error) {
    ctx := context.Background()
    client := openai.NewClient(key)
    model := openai.GPT3Dot5Turbo
    messages := []openai.ChatCompletionMessage{
        {
            Role:    "system",
            Content: "Tugas Anda memberika rekomendasi laptop, cukup berikan infomasi singkat harga yang sesuai dan spesifikasi Merk,Processor,Graphics,RAM,Display,TipePenyimanan",
        },
        {
            Role:    "user",
            Content: fmt.Sprintf("Rekomendasi Laptop dengan budget RP%s untuk %s. dengan merk %s", request.Budget, request.Purpose, request.Merk),
        },
    }

    response, err := us.getCompletionMessages(ctx, client, messages, model)
    if err != nil {
        return "", err
    }
    answer := response.Choices[0].Message.Content
    return answer, nil
}
