import { Injectable } from '@angular/core';
import { HttpClient, HttpResponse } from '@angular/common/http';
import { Todo } from '../models';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root',
})
export class TodoService {
  private backendUrl = environment.backendUrl + '/api/v1';

  constructor(private http: HttpClient) {}

  addTodo(todo: Todo): Observable<HttpResponse<Todo>> {
    return this.http.post<Todo>(this.backendUrl + '/todo', todo, {
      withCredentials: true,
      observe: 'response',
    });
  }

  removeTodo(id: string): Observable<HttpResponse<string>> {
    return this.http.delete<string>(this.backendUrl + '/todo/' + id, {
      withCredentials: true,
      observe: 'response',
    });
  }

  getTodos(): Observable<HttpResponse<Todo[]>> {
    return this.http.get<Todo[]>(this.backendUrl + '/todos', {
      withCredentials: true,
      observe: 'response',
    });
  }

  updateTodo(todo: Todo): Observable<HttpResponse<Todo>> {
    return this.http.put<Todo>(this.backendUrl + '/todo/' + todo.id, todo, {
      withCredentials: true,
      observe: 'response',
    });
  }
}
