# main function goes here
# some git commands:
# git push --force/--force-with-lease to force changes
# git status --short to get short message about current state
# git commit -a to commit changes and skipping staging
# git reset (--soft --mixed --hard) <hash of A>   resets current commit to old commit
# In general, soft: stage everything, mixed: unstage everything, hard: ignore everything up to the commit I'm resetting from

def destroy_walls(wall_healths):
    for wall_health in wall_healths:
        if wall_health <= 0:
            wall_healths.remove(wall_health)
    return wall_healths
