import { Component, OnInit } from '@angular/core';
import { Todo } from '../../models';
import { TodoService } from '../../services/todo.service';
import { Router } from '@angular/router';
import { HttpErrorResponse, HttpResponse } from '@angular/common/http';
import { AuthService } from '../../services/auth.service';

@Component({
  selector: 'app-todo-view',
  templateUrl: './todo-view.component.html',
  styleUrls: ['./todo-view.component.scss'],
})
export class TodoViewComponent implements OnInit {
  todos: Todo[] = [];
  newTodo: string = '';

  constructor(
    private authService: AuthService,
    private todoService: TodoService,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.authService.checkAuth().subscribe({
      next: () => {
        this.authService.isAuthenticated.next(true);
        this.getTodos();
      },
      error: (err) => {
        const error = err as HttpErrorResponse;
        if (error.status === 401) {
          this.authService.isAuthenticated.next(false);
          this.router.navigate(['/', 'login']).then();
        }
      },
    });
  }

  getTodos(): void {
    this.todoService.getTodos().subscribe({
      next: (res) => {
        this.todos = res.body !== null ? res.body : [];
      },
      error: (err) => console.log(err),
    });
  }

  addNewTodo(): void {
    const todo: Todo = {
      id: '',
      name: this.newTodo,
      completed: false,
    };
    if (this.newTodo.length > 0) {
      this.todoService.addTodo(todo).subscribe({
        next: () => (this.newTodo = ''),
        error: (err) => console.log(err),
        complete: () => {
          this.getTodos();
        },
      });
    }
    return;
  }

  deleteTodo(todo: Todo): void {
    this.todoService.removeTodo(todo.id).subscribe({
      next: () => this.getTodos(),
      error: (err) => console.log(err),
    });
  }

  markTodoAsCompleted(todo: Todo): void {
    todo.completed = true;
    this.todoService.updateTodo(todo).subscribe({
      next: () => this.getTodos(),
      error: (err) => console.log(err),
    });
  }
}
