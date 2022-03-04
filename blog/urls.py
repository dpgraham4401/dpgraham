from . import views
from django.urls import path

urlpatterns = [
    path('', views.BlogPostList.as_view(), name='blog_home'),
    path('<slug:slug>', views.BlogPostDetail.as_view(), name='blog_post_detail'),
]
