# from django.shortcuts import render

# Create your views here.
from django.views import generic
from .models import BlogPost

class BlogPostList(generic.ListView):
    queryset = BlogPost.objects.filter(status=1).order_by('-created_date')
    template_name = 'blog_home.html'

class BlogPostDetail(generic.DetailView):
    model = BlogPost
    template_name = 'blog.html'
