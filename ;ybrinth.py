def solution(n, m, obstacles, teleports):
    from collections import deque
    
    obstacle_set = set(map(tuple, obstacles))
    teleport_dict = {}
    for tp in teleports:
        teleport_dict[tuple(tp[:2])] = tuple(tp[2:])
    
    # FIX: Use frozenset() instead of set() for tp_used
    queue = deque([(0, 0, 1, frozenset())])
    # FIX: Change visited to dict to track (position, teleports_used) states
    visited = {}
    visited[(0, 0, frozenset())] = True
    
    while queue:
        row, col, count, tp_used = queue.popleft()
        
        if (row, col) == (n-1, m-1):
            return count
        
        if (row, col) in teleport_dict:
            if (row, col) in tp_used:
                return -2
            # FIX: Use frozenset union operator |
            new_tp_used = tp_used | {(row, col)}
            new_row, new_col = teleport_dict[(row, col)]
            if 0 <= new_row < n and 0 <= new_col < m and (new_row, new_col) not in obstacle_set:
                # FIX: Track state as (position, teleports_used) tuple
                state = (new_row, new_col, new_tp_used)
                if state not in visited:
                    visited[state] = True
                    queue.append((new_row, new_col, count+1, new_tp_used))
            continue
        
        next_col = col+1
        if next_col < m and (row, next_col) not in obstacle_set:
            # FIX: Track state as (position, teleports_used) tuple
            state = (row, next_col, tp_used)
            if state not in visited:
                visited[state] = True
                queue.append((row, next_col, count+1, tp_used))
        else:
            next_row = row+1
            if next_row < n and (next_row, col) not in obstacle_set:
                # FIX: Track state as (position, teleports_used) tuple
                state = (next_row, col, tp_used)
                if state not in visited:
                    visited[state] = True
                    queue.append((next_row, col, count+1, tp_used))
    
    return -1