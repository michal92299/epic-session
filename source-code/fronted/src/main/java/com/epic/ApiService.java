package com.epic;

import com.epic.model.Game;
import com.epic.model.Proton;
import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;

import java.io.IOException;
import java.lang.reflect.Type;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.ArrayList;
import java.util.List;

public class ApiService {

    private static final String API_BASE = "http://localhost:8080";
    private final HttpClient client = HttpClient.newHttpClient();
    private final Gson gson = new Gson();

    public void login() throws IOException, InterruptedException {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(API_BASE + "/login"))
                .GET()
                .build();
        client.send(request, HttpResponse.BodyHandlers.ofString());
    }

    public List<Game> listGames() throws IOException, InterruptedException {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(API_BASE + "/list-games"))
                .GET()
                .build();
        HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());
        Type listType = new TypeToken<ArrayList<Game>>() {}.getType();
        return gson.fromJson(response.body(), listType);
    }

    public void installGame(String appName) throws IOException, InterruptedException {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(API_BASE + "/install-game"))
                .POST(HttpRequest.BodyPublishers.ofString("app_name=" + appName))
                .header("Content-Type", "application/x-www-form-urlencoded")
                .build();
        client.send(request, HttpResponse.BodyHandlers.ofString());
    }

    public void uninstallGame(String appName) throws IOException, InterruptedException {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(API_BASE + "/uninstall-game"))
                .POST(HttpRequest.BodyPublishers.ofString("app_name=" + appName))
                .header("Content-Type", "application/x-www-form-urlencoded")
                .build();
        client.send(request, HttpResponse.BodyHandlers.ofString());
    }

    public List<Proton> listProtons() throws IOException, InterruptedException {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(API_BASE + "/list-protons"))
                .GET()
                .build();
        HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());
        Type listType = new TypeToken<ArrayList<Proton>>() {}.getType();
        return gson.fromJson(response.body(), listType);
    }

    public void installProton(String version) throws IOException, InterruptedException {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(API_BASE + "/install-proton"))
                .POST(HttpRequest.BodyPublishers.ofString("version=" + version))
                .header("Content-Type", "application/x-www-form-urlencoded")
                .build();
        client.send(request, HttpResponse.BodyHandlers.ofString());
    }

    public void launchGame(String appName, String protonPath) throws IOException, InterruptedException {
        String body = "app_name=" + appName + "&proton_path=" + protonPath;
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(API_BASE + "/launch-game"))
                .POST(HttpRequest.BodyPublishers.ofString(body))
                .header("Content-Type", "application/x-www-form-urlencoded")
                .build();
        client.send(request, HttpResponse.BodyHandlers.ofString());
    }

    public String getStatus() throws IOException, InterruptedException {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(API_BASE + "/status"))
                .GET()
                .build();
        HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());
        return response.body();
    }
}
