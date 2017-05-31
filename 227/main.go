package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/shurcooL/play/227/githubql"
	"golang.org/x/oauth2"
)

func main() {
	flag.Parse()

	err := run()
	if err != nil {
		log.Println(err)
	}
}

func run() error {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_GRAPHQL_TEST_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubql.NewClient(httpClient)

	type githubqlActor struct {
		Login     githubql.String
		AvatarURL githubql.URI // TODO: Try out passing arguments, now that I know how.
		URL       githubql.URI
	}

	var v struct {
		Repository struct {
			DatabaseID githubql.Int
			URL        githubql.URI

			Issue struct {
				Author         githubqlActor
				PublishedAt    githubql.DateTime
				LastEditedAt   *githubql.DateTime
				Editor         *githubqlActor
				Body           githubql.String
				ReactionGroups []struct {
					Content githubql.ReactionContent
					Users   struct {
						Nodes []struct {
							Login githubql.String
						}

						TotalCount githubql.Int
					} `graphql:"users(first:10)"`
					ViewerHasReacted githubql.Boolean
				}
				ViewerCanUpdate githubql.Boolean

				Comments struct {
					Nodes []struct {
						Body   githubql.String
						Author struct {
							Login githubql.String
						}
						Editor struct {
							Login githubql.String
						}
					}
					PageInfo struct {
						EndCursor   githubql.String
						HasNextPage githubql.Boolean
					}
				} `graphql:"comments(first:$CommentsFirst,after:$CommentsAfter)"`
			} `graphql:"issue(number:$IssueNumber)"`
		} `graphql:"repository(owner:$RepositoryOwner,name:$RepositoryName)"`
		Viewer struct {
			Login      githubql.String
			CreatedAt  githubql.DateTime
			ID         githubql.ID
			DatabaseID githubql.Int
		}
		RateLimit struct {
			Cost      githubql.Int
			Limit     githubql.Int
			Remaining githubql.Int
			ResetAt   githubql.DateTime
		}
	}
	variables := map[string]interface{}{
		"RepositoryOwner": githubql.String("shurcooL-test"),
		"RepositoryName":  githubql.String("test-repo"),
		"IssueNumber":     githubql.Int(1),
		"CommentsFirst":   githubql.NewInt(1),
		"CommentsAfter":   githubql.NewString("Y3Vyc29yOjE5NTE4NDI1Ng=="),
	}
	err := client.Query(context.Background(), &v, variables)
	if err != nil {
		return err
	}
	printJSON(v)
	//goon.Dump(out)
	//fmt.Println(github.Stringify(out))

	{
		var v struct {
			Repository struct {
				Issue struct {
					ID        githubql.ID
					Reactions struct {
						ViewerHasReacted githubql.Boolean
					} `graphql:"reactions(content:$ReactionContent)"`
				} `graphql:"issue(number:$IssueNumber)"`
			} `graphql:"repository(owner:$RepositoryOwner,name:$RepositoryName)"`
		}
		variables := map[string]interface{}{
			"RepositoryOwner": githubql.String("shurcooL-test"),
			"RepositoryName":  githubql.String("test-repo"),
			"IssueNumber":     githubql.Int(2),
			"ReactionContent": githubql.ReactionContent(githubql.ThumbsUp),
		}
		err := client.Query(context.Background(), &v, variables)
		if err != nil {
			return err
		}
		fmt.Println("already reacted:", v.Repository.Issue.Reactions.ViewerHasReacted)

		if !v.Repository.Issue.Reactions.ViewerHasReacted {
			// Add reaction.
			var m struct {
				AddReaction struct {
					Subject struct {
						ReactionGroups []struct {
							Users struct {
								TotalCount githubql.Int
							}
						}
					}
				} `graphql:"addReaction(input:$Input)"`
			}
			variables := map[string]interface{}{
				"Input": githubql.AddReactionInput{
					SubjectID: v.Repository.Issue.ID,
					Content:   githubql.ThumbsUp,
				},
			}
			err := client.Mutate(context.Background(), &m, variables)
			if err != nil {
				return err
			}
			printJSON(m)
		} else {
			// TODO: Use GraphQL API.
		}
	}

	return nil
}

// printJSON prints v as JSON encoded with indent to stdout. It panics on any error.
func printJSON(v interface{}) {
	w := json.NewEncoder(os.Stdout)
	w.SetIndent("", "\t")
	err := w.Encode(v)
	if err != nil {
		panic(err)
	}
}
