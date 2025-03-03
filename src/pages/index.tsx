"use client"

import { useState, useEffect, useCallback } from "react"
import { Loader2, RefreshCw } from "lucide-react"
import { Button } from "@/components/button"
import { Input } from "@/components/input"
import { Textarea } from "@/components/textarea"

// Mock post type
type Post = {
  content: string;
  created_at: string;
  status_id: string;
};

export default function Home() {
  // State for post creation
  const [postContent, setPostContent] = useState("")
  const [isCreatingPost, setIsCreatingPost] = useState(false)

  // State for post deletion
  const [postIdToDelete, setPostIdToDelete] = useState("")

  // State for posts list
  const [posts, setPosts] = useState<Post[]>([])
  const [isLoading, setIsLoading] = useState(false)
  const [filter, setFilter] = useState("all")

  // Character limit
  const maxCharacters = 400

  // Function to strip HTML tags
  const stripHTMLTags = (str: string) => str.replace(/<[^>]*>/g, '');


  // Create post handler
  const handleCreatePost = async () => {
    if (!postContent.trim()) return
  
    setIsCreatingPost(true)
  
    try {
      const response = await fetch('http://192.168.12.186:8080/posts', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          status: postContent,
        }),
      })
  
      if (!response.ok) {
        throw new Error('Failed to create post')
      }
  
      const newPost: Post = await response.json()
      console.log(newPost)
      setPostContent("")
    } catch (error) {
      console.error('Error creating post:', error)
      // Handle error (e.g., show error message to user)
    } finally {
      setIsCreatingPost(false)
    }
  }  

  // Delete post handler
  const handleDeletePost = async () => {
    if (!postIdToDelete) return
  
    try {
      const response = await fetch(`http://192.168.12.186:8080/posts/${postIdToDelete}`, {
        method: 'DELETE',
      })
  
      if (!response.ok) {
        throw new Error('Failed to delete post')
      }
  
      // If the deletion was successful, update the local state
      setPosts((prev) => prev.filter((post) => post.status_id !== postIdToDelete))
      setPostIdToDelete("")
    } catch (error) {
      console.error('Error deleting post:', error)
      // Handle error (e.g., show error message to user)
    }
  }
  

  // Refresh posts handler
  const handleRefreshPosts = useCallback(async () => {
    setIsLoading(true);
  
    try {
      const response = await fetch('http://192.168.12.186:8080/posts', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
  
      if (!response.ok) {
        throw new Error('Failed to fetch posts');
      }
  
      const result = await response.json();
      const fetchedPosts: Post[] = result.data;
  
      setPosts(fetchedPosts);
    } catch (error) {
      console.error('Error fetching posts:', error);
      // Handle error (e.g., show error message to user)
    } finally {
      setIsLoading(false);
    }
  }, []);
  
  

  // Load initial posts
  useEffect(() => {
    handleRefreshPosts()
  }, [handleRefreshPosts])

  // Filter posts based on selection
  const filteredPosts =
    filter === "all"
      ? posts
      : filter === "recent"
        ? [...posts].sort((a, b) => new Date(b.created_at).getTime() - new Date(a.created_at).getTime())
        : posts.slice(0, 2) // 'limited' filter just shows first 2 posts

  return (
    <main className="container mx-auto py-8 px-4">
      <h1 className="text-3xl font-bold mb-8 text-center">Team AlphaGO Twitter Service  </h1>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
        {/* Left Column - Post Creation and Deletion */}
        <div className="space-y-8">
          <div className="space-y-4 p-6 border rounded-lg shadow-sm">
            <h2 className="text-xl font-semibold">Create New Post</h2>

            <div className="space-y-2 w-full">
              <Textarea
                placeholder="Write your post content here..."
                value={postContent}
                onChange={(e) => setPostContent(e.target.value.slice(0, maxCharacters))}
                className="min-h-[150px] resize-none w-full"
              />
              <div className="text-sm text-muted-foreground text-right">
                {postContent.length}/{maxCharacters} characters
              </div>
            </div>

            <Button onClick={handleCreatePost} disabled={!postContent.trim() || isCreatingPost} className="w-full">
              {isCreatingPost ? (
                <>
                  <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                  Creating...
                </>
              ) : (
                "Create Post"
              )}
            </Button>
          </div>

          <div className="space-y-4 p-6 border rounded-lg shadow-sm">
            <h2 className="text-xl font-semibold">Delete Post</h2>

            <div className="space-y-4 w-full">
              <Input className="w-full"
                placeholder="Enter post ID to delete"
                value={postIdToDelete}
                onChange={(e) => setPostIdToDelete(e.target.value)}
              />

              <Button onClick={handleDeletePost} disabled={!postIdToDelete} variant="destructive" className="w-full">
                Delete Post
              </Button>
            </div>
          </div>
        </div>

        {/* Right Column - Posts List */}
        <div className="space-y-4">
          <div className="flex items-center justify-between">
            <div className="flex items-center space-x-2">
              <h2 className="text-xl font-semibold">Posts</h2>
              <Button variant="outline" size="icon" onClick={handleRefreshPosts} disabled={isLoading}>
                <RefreshCw className={`h-4 w-4 ${isLoading ? "animate-spin" : ""}`} />
                <span className="sr-only">Refresh posts</span>
              </Button>
            </div>
          </div>

          {isLoading ? (
            <div className="flex justify-center items-center h-64">
              <Loader2 className="h-8 w-8 animate-spin text-primary" />
            </div>
          ) : filteredPosts.length > 0 ? (
            <div className="space-y-4 overflow-auto max-h-[600px] pr-2">
              {filteredPosts.map((post) => (
                <div key={post.status_id} className="p-4 border rounded-lg shadow-sm">
                  <div className="flex justify-between items-start mb-2">
                    <span className="text-sm font-medium bg-primary/10 text-primary px-2 py-1 rounded">
                      ID: {post.status_id}
                    </span>
                    <span className="text-xs text-muted-foreground">{new Date(post.created_at).toLocaleString()}</span>
                  </div>
                  <p className="text-sm">{stripHTMLTags(post.content)}</p>
                </div>
              ))}
            </div>
          ) : (
            <div className="flex justify-center items-center h-64 border rounded-lg">
              <p className="text-muted-foreground">No posts found</p>
            </div>
          )}
        </div>
      </div>
    </main>
  )
}

