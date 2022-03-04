from django.contrib import admin
from .models import BlogPost

class BlogPostAdmin(admin.ModelAdmin):
    list_display = ('title', 'slug', 'status','created_date')
    list_filter = ("status",)
    search_fields = ['title', 'content']

# Register your models here.
admin.site.register(BlogPost, BlogPostAdmin)
